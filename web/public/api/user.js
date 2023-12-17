import { instance } from "./config.js";
import useSWR, { mutate } from "swr";
import { useState, useEffect } from "react";
import { useSession } from "../hooks/session.js";

/**
 * createUser creates a new user account.
 */
export const createUser = (data, fetcher) => {
	return fetcher(`${instance}/api/v1/users`, {
		body: JSON.stringify(data),
		method: "POST",
	}).then((response) => {
		mutate(`${instance}/api/v1/users`);
		return response;
	});
};

/**
 * deleteUser deletes a named user
 */
export const deleteUser = (data, fetcher) => {
	const { id } = data;
	return fetcher(`${instance}/api/v1/users/${id}`, {
		method: "DELETE",
	}).then((response) => {
		mutate(`${instance}/api/v1/users`);
		return response;
	});
};

/**
 * updateCurrentUser updates the currently authenticated user.
 */
export const updateCurrentUser = (data, fetcher) => {
	return fetcher(`${instance}/api/v1/user`, {
		body: JSON.stringify(data),
		method: "PATCH",
	}).then((response) => {
		mutate(`${instance}/api/v1/user`);
		return response;
	});
};

/**
 * registerUser registers a new user account.
 */
export const registerUser = (data, fetcher) => {
	return fetcher(`${instance}/api/v1/register`, {
		body: data,
		method: "POST",
	});
};

/**
 * authenticateUser authenticates a user account.
 */
export const authenticateUser = (data, fetcher) => {
	return fetcher(`${instance}/api/v1/login`, {
		body: data,
		method: "POST",
	});
};

/**
 * useToken generates a user token.
 */
export const useToken = (fetcher) => {
	const [data, setData] = useState("");
	const [error, setError] = useState(null); // TODO(bradrydzewski) setError

	const fetchToken = async () => {
		fetcher(`${instance}/api/v1/user/token`, {
			credentials: "same-origin",
			method: "POST",
		}).then((token) => {
			setData(token);
		});
	};

	useEffect(() => {
		fetchToken();
	}, []);

	return {
		token: data,
		isLoading: !error && !data,
		isError: error,
	};
};

/**
 * useUserList returns an swr hook that provides a list of users.
 */
export const useUserList = () => {
	const { data, error } = useSWR(`${instance}/api/v1/users`);
	return {
		userList: data,
		isLoading: !error && !data,
		isError: error,
	};
};

/**
 * useUser returns an swr hook that provides a user.
 */
export const useUser = (id) => {
	const { data, error } = useSWR(`${instance}/api/v1/users/${id}`);
	return {
		user: data,
		isLoading: !error && !data,
		isError: error,
	};
};
