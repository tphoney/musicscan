import styles from "./account.module.css";
import { instance } from "../api/config.js";
import { useToken } from "../api/user.js";
import { useSession } from "../hooks/session.js";

// address provides the remote API address.
const address =
	instance || `${window.location.protocol}//${window.location.host}`;

// Renders the Account page.
export default function Account({ params }) {
	const { session, signout, fetcher } = useSession();
	const { token } = useToken(fetcher);
	return (
		<>
			<section className={styles.root}>
				<h2>Token ({session && session.user.email})</h2>
				<pre>{token && token.access_token}</pre>
				<pre>{`curl -H "Authorization: Bearer ${
					token && token.access_token
				}" ${address}/api/v1/user`}</pre>
				<button onClick={signout}>Logout</button>
			</section>
		</>
	);
}
