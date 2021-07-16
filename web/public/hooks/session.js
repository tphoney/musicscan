import { useState, useContext, createContext } from "react";

const initialState = (function () {
	const data = localStorage.getItem("session");
	return data && JSON.parse(data);
})();

const sessionContext = createContext(null);

export function ProvideSession({ children }) {
	const session = useProvideSession();
	return (
		<sessionContext.Provider value={session}>
			{children}
		</sessionContext.Provider>
	);
}

export const useSession = () => {
	return useContext(sessionContext);
};

function useProvideSession() {
	const [session, setSession] = useState(initialState);

	// signin adds the session to the context and
	// to local storage.
	const signin = (session) => {
		localStorage.setItem("session", JSON.stringify(session));
		setSession(session);
	};

	// signout removes the session from the context
	// and from local storage.
	const signout = () => {
		localStorage.removeItem("session");
		setSession(null);
	};

	// update updates the user account details stored
	// in the current session.
	const update = (session) => {
		localStorage.setItem("session", JSON.stringify(session));
		setSession(session);
	};

	const fetcher = async (endpoint, initConfig) => {
		const credentials = "same-origin";
		const headers = new Headers(
			session && session.token
				? { Authorization: `Bearer ${session.token.access_token}` }
				: {}
		);
		const res = await fetch(endpoint, { ...initConfig, headers, credentials });

		// if the response status is 401 it indicates the
		// session token is expired. redirect the user to
		// the login screen.
		if (res.status === 401) {
			signout();
			return res;
		}

		// if the response succeeds but returns no content
		// return the response. do not attempt to unmarshal
		// the response body, since none exists.
		if (res.status === 204) {
			return res;
		}

		// if the reponse succeeds we can assume the response
		// is in json format (since the server returns all
		// responses in json format) so we can unmarshal the
		// body and return the object.
		if (res.ok) {
			return res.json();
		}

		// if an error response is returned in json format,
		// unmarshal and return the object as an error.
		const contentType = res.headers.get("content-type") || "";
		if (contentType.startsWith("application/json")) {
			return res.json().then((error) => {
				throw error;
			});
		}

		// else return the response body text as the error message.
		return res.text().then((error) => {
			throw error;
		});
	};

	return {
		session,
		signin,
		signout,
		fetcher,
		update,
	};
}
