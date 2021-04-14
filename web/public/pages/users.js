import { useState, useRef } from "react";
import styles from "./users.module.css";
import { useUserList, createUser, deleteUser } from "../api/user.js";
import { useSession } from "../hooks/session.js";

// Renders the User List page.
export default function Users({ params }) {
	const { fetcher } = useSession();

	//
	// Load User List
	//

	const { userList, isLoading, isError } = useUserList();
	if (isLoading) {
		return renderLoading();
	}
	if (isError) {
		return renderError(isError);
	}
	if (userList.length === 0) {
		return renderEmpty();
	}

	//
	// Add User Functions
	//

	const [error, setError] = useState(null);
	const emailElem = useRef(null);
	const roleElem = useRef(null);

	const handleCreate = () => {
		const email = emailElem.current.value;
		const admin = roleElem.current.value === "admin";
		createUser({ email, admin }, fetcher)
			.then((user) => {
				emailElem.current.value = "";
			})
			.catch((error) => {
				console.log(error);
			});
	};

	//
	// Delete Functions
	//

	const handleDelete = (user) => {
		deleteUser(user, fetcher);
	};

	//
	// Render the Page
	//

	return (
		<>
			<section className={styles.root}>
				<h2>Users</h2>
				<ul>
					{userList.map((user) => (
						<UserInfo user={user} onDelete={handleDelete} />
					))}
				</ul>

				<div className="actions">
					<button onClick={handleCreate}>Add User</button>
					<input ref={emailElem} type="text" placeholder="email" />
					<select ref={roleElem}>
						<option value="developer" selected>
							developer
						</option>
						<option value="admin">admin</option>
					</select>
				</div>
			</section>
		</>
	);
}

// render the user information.
const UserInfo = ({ user, onDelete }) => {
	return (
		<li id={user.id}>
			{user.email}
			<button onClick={onDelete.bind(this, user)}>Delete</button>
		</li>
	);
};

// helper function renders the loading bar.
const renderLoading = () => {
	return <div>Loading ...</div>;
};

// helper function returns the error message.
const renderError = (error) => {
	return <div>{error}</div>;
};

// helper function returns the empty message.
const renderEmpty = (error) => {
	return <div>Your Artist list is empty</div>;
};
