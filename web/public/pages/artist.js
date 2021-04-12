import styles from "./project.module.css";
import { Route, Switch } from "wouter";
import { useartist } from "../api/artist.js";
import { useProject } from "../api/project.js";

import Link from "../shared/link.js";

import albumList from "./album_list.js";

// Renders the artist page.
export default function artist({ params }) {
	//
	// Load Project
	//

	const {
		project,
		isLoading: isProjectLoading,
		isError: isProjectError,
	} = useProject(params.project);

	if (isProjectLoading) {
		return renderLoading();
	}
	if (isProjectError) {
		return renderError(isProjectError);
	}

	//
	// Load artist
	//

	const { artist, isLoading: isartistLoading, isError: isartistErrror } = useartist(
		params.project,
		params.artist
	);

	if (isartistLoading) {
		return renderLoading();
	}
	if (isartistErrror) {
		return renderError(isartistErrror);
	}

	//
	// Render Page
	//

	return (
		<>
			<nav>
				<h1>{artist.name}</h1>
				<ul>
					<li>
						<Link href={`/projects/${project.id}/artists/${artist.id}`}>
							albums
						</Link>
					</li>
					<li>
						<Link href="#">Details</Link>
					</li>
				</ul>
			</nav>

			<Switch>
				<Route path="/projects/:project/artists/:artist" component={albumList} />
				<Route>Not Found</Route>
			</Switch>
		</>
	);
}

// helper function renders the loading bar.
const renderLoading = () => {
	return <div>Loading ...</div>;
};

// helper function returns the error message.
const renderError = (error) => {
	return <div>{error}</div>;
};
