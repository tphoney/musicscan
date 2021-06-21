import { useState, useRef } from "react";
import styles from "./account.module.css";
import { instance } from "../../api/config.js";
import { useToken, updateCurrentUser } from "../../api/user.js";
import { useSession } from "../../hooks/session.js";
import { useProjectList } from "../../api/project.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";

// address provides the remote API address.
const address =
	instance || `${window.location.protocol}//${window.location.host}`;

// Renders the Account page.
export default function Account({ params }) {
	const { session, fetcher, update } = useSession();
	const { token } = useToken(fetcher);

	const [showToken, setShowToken] = useState(false);
	const [userData, setUserData] = useState({
		email: session.user.email,
		name: session.user.name,
		company: session.user.company,
	});

	const passwordElem1 = useRef(null);
	const passwordElem2 = useRef(null);

	//
	// Load Project List
	//

	const { projectList } = useProjectList();

	//
	// Handle Update User
	//

	const handleUpdateEmail = (event) => {
		setUserData({
			email: event.target.value,
			name: userData.name,
			company: userData.company,
		});
	};

	const handleUpdateName = (event) => {
		setUserData({
			name: event.target.value,
			email: userData.email,
			company: userData.company,
		});
	};

	const handleUpdateCompany = (event) => {
		setUserData({
			email: userData.email,
			name: userData.name,
			company: event.target.value,
		});
	};

	const handleUpdate = () => {
		updateCurrentUser(userData, fetcher).then(() => {
			session.user.name = userData.name;
			session.user.email = userData.email;
			session.user.company = userData.company;
			update(session);
		});
	};

	const handleUpdatePassword = () => {
		const password1 = passwordElem1.current.value;
		const password2 = passwordElem2.current.value;
		if (password1 !== password2) {
			window.alert("Passwords do not match");
			return;
		}
		updateCurrentUser({ password: password1 }, fetcher).then(() => {
			passwordElem1.current.value = "";
			passwordElem2.current.value = "";
		});
	};

	const handleDeleteAccount = () => {
		window.alert("not implemented");
	};

	return (
		<>
			<section className={styles.root}>
				<div className={styles.card}>
					<h2>Profile</h2>
					<div className={styles.field}>
						<label>Email *</label>
						<Input
							type="text"
							value={userData.email}
							onChange={handleUpdateEmail}
						/>
					</div>
					<div className={styles.field}>
						<label>Full Name *</label>
						<Input
							type="text"
							value={userData.name}
							onChange={handleUpdateName}
						/>
					</div>
					<div className={styles.field}>
						<label>Company</label>
						<Input
							type="text"
							value={userData.company}
							onChange={handleUpdateCompany}
						/>
					</div>
					<div className={styles.actions}>
						<Button onClick={handleUpdate}>Update Profile</Button>
					</div>
				</div>

				<div className={styles.card}>
					<h2>Credentials</h2>
					<div className={styles.field}>
						<label>Password</label>
						<Input type="password" ref={passwordElem1} />
					</div>
					<div className={styles.field}>
						<label>Re-type your Password</label>
						<Input type="password" ref={passwordElem2} />
					</div>
					<div className={styles.actions}>
						<Button onClick={handleUpdatePassword}>Update Password</Button>
					</div>
				</div>

				<div className={styles.card}>
					<h2>Workspaces</h2>
					<div className={styles.workspaces}>
						{projectList &&
							projectList.map((project) => (
								<div>
									<Avatar text={project.name} className={styles.avatar} />
									{project.name}
								</div>
							))}
					</div>
				</div>

				<div className={styles.card}>
					<h2>Token</h2>
					<p>Personal access token that can be used to access the API.</p>
					{showToken && <pre>{token && token.access_token}</pre>}
					{!showToken && (
						<Button onClick={() => setShowToken(true)}>Display Token</Button>
					)}
				</div>

				<div className={styles.card}>
					<h2>Delete Account</h2>
					<p>Warning, this action cannot be undone.</p>
					<Button onClick={handleDeleteAccount}>Delete</Button>
				</div>
			</section>
		</>
	);
}
