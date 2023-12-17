import styles from "./artist.module.css";
import { Route, Switch } from "wouter";
import { useArtist } from "../../api/artist.js";
import { useProject } from "../../api/project.js";

import AlbumList from "../albums/albums.js";
import AlbumInfo from "../album/album.js";

// Renders the Artist page.
export default function Artist({ params }) {
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
	// Load Artist
	//

	const { artist, isLoading: isArtistLoading, isError: isArtistErrror } = useArtist(
		params.project,
		params.artist
	);

	if (isArtistLoading) {
		return renderLoading();
	}
	if (isArtistErrror) {
		return renderError(isArtistErrror);
	}

	//
	// Render Page
	//

	return (
		<>
			<Switch>
				<Route
					path="/projects/:project/artists/:artist/albums/:album"
					component={AlbumInfo}
				/>
				<Route
					path="/projects/:project/artists/:artist"
					component={AlbumList}
				/>
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
