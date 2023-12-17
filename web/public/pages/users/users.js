import { useState, useRef } from "react";
import styles from "./users.module.css";
import { useUserList, createUser, deleteUser } from "../../api/user.js";
import { useSession } from "../../hooks/session.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";
import Select from "../../shared/components/select";

import { Drawer, Target } from "@accessible/drawer";

// Renders the User List page.
export default function Users({ params }) {
	const { fetcher } = useSession();
	const [open, setOpen] = useState(false);

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
	const passElem = useRef(null);

	const handleCreate = () => {
		const email = emailElem.current.value;
		const admin = roleElem.current.value === "admin";
		const password = passElem.current.value;
		createUser({ email, admin, password }, fetcher)
			.then((user) => {
				emailElem.current.value = "";
				passElem.current.value = "";
				setOpen(false);
			})
			.catch((error) => {
				console.log(error);
			});
	};

	//
	// Delete Functions
	//

	const handleDelete = (user) => {
		if (confirm("Are you sure you want to proceed?")) {
			deleteUser(user, fetcher);
		}
	};

	//
	// Render the Page
	//

	return (
		<>
			<section className={styles.root}>
				<p>Manage the list of users that have access to the system.</p>

				<ul className={styles.list}>
					{userList.map((user) => (
						<UserInfo user={user} onDelete={handleDelete} />
					))}
				</ul>
				<Button className={styles.button} onClick={() => setOpen(true)}>
					New User
				</Button>
			</section>

			<Drawer open={open}>
				<Target
					placement="right"
					closeOnEscape={true}
					preventScroll={true}
					openClass={styles.drawer}
				>
					<div>
						<Input ref={emailElem} type="text" placeholder="email" />
						<Input ref={passElem} type="password" placeholder="password" />
						<Select ref={roleElem}>
							<option value="developer" selected>
								developer
							</option>
							<option value="admin">admin</option>
						</Select>

						<div className={styles.actions}>
							<Button onClick={handleCreate}>Add User</Button>
							<Button onClick={() => setOpen(false)}>Close</Button>
						</div>
					</div>
				</Target>
			</Drawer>
		</>
	);
}

// render the user information.
const UserInfo = ({ user, onDelete }) => {
	return (
		<li id={user.id} className={styles.item}>
			<Avatar text={user.email} className={styles.avatar} />
			<span className={styles.fill}>
				{user.email} ({user.admin ? "admin" : "developer"})
			</span>
			<Button onClick={onDelete.bind(this, user)}>Delete</Button>
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
