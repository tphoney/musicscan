import styles from "./project.module.css";
import { Route, Switch } from "wouter";
import { useProject } from "../api/project.js";

import Link from "../shared/link.js";

import Artist from "./artist.js";
import ArtistList from "./artist_list.js";
import Member from "./members.js";
import Scan from "./project_scan.js";
import ProjectBadAlbumList from "./project_bad_album_list.js";

// Renders the Project page.
export default function Project({ params }) {
	const { project, isLoading, isError } = useProject(params.project);

	if (isLoading) {
		return renderLoading();
	}
	if (isError) {
		return renderError(isError);
	}

	return (
		<>
			<nav>
				<h1>{project && project.name}</h1>
				<ul>
					<li>
						<Link href={`/projects/${project.id}`}>Artists</Link>
					</li>
					<li>
						<Link href={`/projects/${project.id}/members`}>Members</Link>
					</li>
					<li>
						<Link href={`/projects/${project.id}/project_scan`}>Scan</Link>
					</li>
					<li>
						<Link href={`/projects/${project.id}/project_bad_album_list`}>Bad Albums</Link>
					</li>
				</ul>
			</nav>

			<Switch>
				<Route path="/projects/:project" component={ArtistList} />
				<Route path="/projects/:project/artists" component={ArtistList} />
				<Route path="/projects/:project/artists/:artist" component={Artist} />
				<Route path="/projects/:project/artists/path+" component={Artist} />
				<Route path="/projects/:project/members" component={Member} />
				<Route path="/projects/:project/project_scan" component={Scan} />
				<Route path="/projects/:project/project_bad_album_list" component={ProjectBadAlbumList} />
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
