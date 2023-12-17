import styles from "./shell.module.css";
import classnames from "classnames";
import { useSession } from "../../../hooks/session.js";

import { Route, Switch, useLocation } from "wouter";

import Project from "./project";

import Link from "../../link.js";
import Avatar from "../../components/avatar";
import Button from "../../components/button";

export default (props) => {
	const { session, signout } = useSession();
	const [location, setLocation] = useLocation();

	if (!isProjectLayout(location)) {
		return (
			<>
				<header className={styles.header}>
					<Link href="/">
						<img src="/logo.svg" />
					</Link>
					<div className={styles.menu}></div>
					{session.user.admin ? (
						<div className={styles.logout}>
							<Button onClick={() => setLocation("/users")}>Admin</Button>
						</div>
					) : undefined}
					{location === "/account" ? (
						<div className={styles.logout}>
							<Button onClick={signout}>Logout</Button>
						</div>
					) : undefined}
					<Link href="/account">
						<Avatar text={session.user.email} />
					</Link>
				</header>

				<div className={classnames(styles.root, props.className)}>
					{props.children}
				</div>
			</>
		);
	}

	return <Project>{props.children}</Project>;
};

// helper function returns true if the page should
// use a project layout.
function isProjectLayout(location) {
	return (
		location != "" &&
		location != "/" &&
		location != "/account" &&
		location != "/users" &&
		!location.startsWith("/account/") &&
		!location.startsWith("/users/")
	);
}
