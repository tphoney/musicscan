import { useRef, useState } from "react";
import styles from "./login.module.css";
import { authenticateUser } from "../../api/user.js";
import { useSession } from "../../hooks/session.js";

import Link from "../../shared/link.js";
import Input from "../../shared/components/input";
import Button from "../../shared/components/button";

// Renders the Login page.
export default function Login({ params }) {
	const { signin, fetcher } = useSession();

	const [error, setError] = useState(null);
	const usernameElem = useRef(null);
	const passwordElem = useRef(null);

	const handleLogin = () => {
		let formData = new FormData();
		formData.append("username", usernameElem.current.value);
		formData.append("password", passwordElem.current.value);
		authenticateUser(formData, fetcher)
			.then((session) => {
				usernameElem.current.value = "";
				passwordElem.current.value = "";
				signin(session);
			})
			.catch((error) => {
				setError(error);
			});
	};

	const alert =
		error && error.message ? (
			<div class="alert">{error.message}</div>
		) : undefined;

	return (
		<>
			<div className={styles.root}>
				<div className={styles.logo}>
					<img src="/logo.svg" />
				</div>
				<h2>Log in with your account</h2>
				{alert}
				<div className={styles.field}>
					<label>Email</label>
					<Input
						ref={usernameElem}
						type="text"
						name="username"
						placeholder="Email"
						className={styles.input}
					/>
				</div>
				<div className={styles.field}>
					<label>Password</label>
					<Input
						ref={passwordElem}
						type="password"
						name="password"
						placeholder="Password"
						className={styles.input}
					/>
				</div>
				<div>
					<Button onClick={handleLogin} className={styles.submit}>
						Log In
					</Button>
				</div>
				<div className={styles.actions}>
					<span>
						Don't have an account? <Link href="/register">Sign Up</Link>
					</span>
				</div>
			</div>
		</>
	);
}
