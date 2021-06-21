import { Route, Switch } from "wouter";
import { useProject } from "../../api/project.js";

import Artist from "../artist/artist.js";
import ArtistList from "../artists/artists.js";
import ArtistSettings from "../artist/settings.js";
import Member from "../members/members.js";
import Settings from "./settings.js";
import Analysis from "../analysis/analysis.js";
import BadAlbumList from "../analysis/badalbums.js";
import WantedAlbumList from "../analysis/wantedalbums.js";

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
			<Switch>
				<Route path="/projects/:project" component={ArtistList} />
				<Route path="/projects/:project/artists" component={ArtistList} />
				<Route path="/projects/:project/artists/:artist" component={Artist} />
				<Route path="/projects/:project/artists/:artist/settings" component={ArtistSettings} />
				<Route path="/projects/:project/artists/:artist/albums/:album" component={Artist} />
				<Route path="/projects/:project/artists/path+" component={Artist} />
				<Route path="/projects/:project/analysis" component={Analysis} />
				<Route path="/projects/:project/analysis/bad_album_list" component={BadAlbumList} />
				<Route path="/projects/:project/analysis/wanted_album_list/:year" component={WantedAlbumList} />
				<Route path="/projects/:project/members" component={Member} />
				<Route path="/projects/:project/settings" component={Settings} />
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
