import { useEffect, useState } from "react";
import { render } from "react-dom";
import { Route, Switch, useRoute } from "wouter";
import { SWRConfig } from "swr";
// import { fetcher } from "./api/config";

import Link from "./shared/link.js";
import { ProvideSession, useSession } from "./hooks/session.js";

import Account from "./pages/account.js";
import Home from "./pages/home.js";
import Login from "./pages/login.js";
import Project from "./pages/project.js";
import Register from "./pages/register.js";
import Users from "./pages/users.js";

export default function App() {
	const { session, fetcher } = useSession();

	// if the session is loaded, and the session
	// is falsey, the login and register routes
	// are exposed.
	if (!session) {
		return (
			<>
				<Switch>
					<Route path="/register" component={Register} />
					<Route component={Login} />
				</Switch>
			</>
		);
	}

	return (
		<>
			<nav>
				<ul>
					<li>
						<Link href="/">Home</Link>
					</li>
					{session.user.admin ? (
						<li>
							<Link href="/users">Users</Link>
						</li>
					) : undefined}
					<li>
						<Link href="/account">Account</Link>
					</li>
				</ul>
			</nav>
			<SWRConfig value={{ fetcher }}>
				<Switch>
					<Route path="/" component={Home} />
					<Route path="/users" component={Users} />
					<Route path="/projects/:project" component={Project} />
					<Route path="/projects/:project/:path+" component={Project} />
					<Route path="/account" component={Account} />
					<Route>Not Found</Route>
				</Switch>
			</SWRConfig>
		</>
	);
}

render(
	<ProvideSession>
		<App />
	</ProvideSession>,
	document.body
);
