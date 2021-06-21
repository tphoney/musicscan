// @ts-nocheck
import styles from "./simple.module.css";
import classnames from "classnames";
import { useSession } from "../../../hooks/session.js";

import { useLocation } from "wouter";

import Link from "../../link.js";
import Avatar from "../../components/avatar";
import Button from "../../components/button";

export default ({ header: Header, content: Content, className, ...rest }) => {
	const { session, signout } = useSession();
	const [location, setLocation] = useLocation();

	return (
		<main className={styles.root}>
			<header className={styles.header}>
				<Link href="/">
					<img src="/logo.svg" />
				</Link>
				<div className={styles.menu}>
					{location === "/account" ? (
						<div className={styles.logout}>
							<Button onClick={signout}>Logout</Button>
						</div>
					) : undefined}
					{session.user.admin ? (
						<div className={styles.logout}>
							<Button onClick={() => setLocation("/users")}>Admin</Button>
						</div>
					) : undefined}
				</div>
				<Link href="/account">
					<Avatar text={session.user.email} />
				</Link>
			</header>

			<div className={classnames(styles.content, className)}>
				<Content {...rest} />
			</div>
		</main>
	);
};
