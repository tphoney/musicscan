import { useRef, useState } from "react";
import { useLocation } from "wouter";
import styles from "./register.module.css";
import { registerUser } from "../../api/user.js";
import { useSession } from "../../hooks/session.js";

import Link from "../../shared/link.js";
import Input from "../../shared/components/input";
import Button from "../../shared/components/button";

// Renders the Register page.
export default function Register({ params }) {
	const { signin, fetcher } = useSession();
	const [location, setLocation] = useLocation();

	const [error, setError] = useState(null);
	const usernameElem = useRef(null);
	const passwordElem = useRef(null);
	const fullnameElem = useRef(null);

	const handleRegister = () => {
		let formData = new FormData();
		formData.append("username", usernameElem.current.value);
		formData.append("password", passwordElem.current.value);
		formData.append("fullname", fullnameElem.current.value);
		registerUser(formData, fetcher)
			.then((session) => {
				usernameElem.current.value = "";
				passwordElem.current.value = "";
				fullnameElem.current.value = "";
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
			<div className={styles.root}>
				<div className={styles.logo}>
					<img src="/logo.svg" />
				</div>
				<h2>Sign up for a new account</h2>
				{alert}
				<div className={styles.field}>
					<label>Full Name</label>
					<Input
						ref={fullnameElem}
						type="text"
						name="fullname"
						placeholder="Full Name"
						className={styles.input}
					/>
				</div>
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
					<Button onClick={handleRegister} className={styles.submit}>
						Sign Up
					</Button>
				</div>
				<div className={styles.actions}>
					<span>
						Already have an account? <Link href="/login">Sign In</Link>
					</span>
				</div>
			</div>
		</>
	);
}
