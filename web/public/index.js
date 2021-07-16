import { render } from "react-dom";
import { Route, Switch, Redirect } from "wouter";
import { SWRConfig } from "swr";

import { ProvideSession, useSession } from "./hooks/session.js";

import Account from "./pages/account/account.js";
import Home from "./pages/projects/projects.js";
import Login from "./pages/login/login.js";
import Project from "./pages/project/project.js";
import Register from "./pages/register/register.js";
import Users from "./pages/users/users.js";

import Shell from "./shared/layouts/shell/shell.js";
import { SimpleRoute, ComplexRoute } from "./shared/router/router";
import Guest from "./shared/layouts/login.js";

// TODO remove me
import Demo from "./shared/components/demo/demo.js";

export default function App() {
	const { session, fetcher } = useSession();

	// if the session is loaded, and the session
	// is falsey, the login and register routes
	// are exposed.
	if (!session) {
		return (
			<>
				<Guest>
					<Switch>
						<Route path="/demo" component={Demo} />
						<Route path="/register" component={Register} />
						<Route component={Login} />
					</Switch>
				</Guest>
			</>
		);
	}

	return (
		<>
			<SWRConfig value={{ fetcher }}>
				<Switch>
					<SimpleRoute path="/" content={Home} />
					<SimpleRoute path="/users" content={Users} />
					<SimpleRoute path="/account" content={Account} />
					<ComplexRoute path="/projects/:project" content={Project} />
					<ComplexRoute path="/projects/:project/:path+" content={Project} />
					<Route path="/login">
						<Redirect to={"/"} />
					</Route>
					<Route path="/register">
						<Redirect to={"/"} />
					</Route>
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
