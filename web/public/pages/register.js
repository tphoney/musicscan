import { useRef, useState } from "react";
import { useLocation } from "wouter";
import styles from "./register.module.css";
import { registerUser } from "../api/user.js";
import { useSession } from "../hooks/session.js";

// Renders the Register page.
export default function Register({ params }) {
	const { signin, fetcher } = useSession();
	const [location, setLocation] = useLocation();

	const [error, setError] = useState(null);
	const usernameElem = useRef(null);
	const passwordElem = useRef(null);

	const handleRegister = () => {
		let formData = new FormData();
		formData.append("username", usernameElem.current.value);
		formData.append("password", passwordElem.current.value);
		registerUser(formData, fetcher)
			.then((session) => {
				usernameElem.current.value = "";
				passwordElem.current.value = "";
				signin(session);
				setLocation("/");
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
			<section className={styles.root}>
				<h2>Register</h2>
				{alert}
				<input
					ref={usernameElem}
					type="text"
					name="username"
					placeholder="Email"
				/>
				<input
					ref={passwordElem}
					type="password"
					name="password"
					placeholder="Password"
				/>
				<button onClick={handleRegister}>Register</button>
			</section>
		</>
	);
}
