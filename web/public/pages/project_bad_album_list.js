import { useState, useRef } from "react";
import styles from "./album_list.module.css";
import { Link } from "wouter";
import { useAlbumBadList } from "../api/project.js";
import { useSession } from "../hooks/session.js";

// Renders the Album List page.
export default function ProjectBadAlbumList({ params }) {
	const { fetcher } = useSession();

	//
	// Load Project
	//

	const {
		badAlbumList,
		isLoading: isProjectLoading,
		isError: isProjectError,
	} = useAlbumBadList(params.project);

	if (isProjectLoading) {
		return renderLoading();
	}
	if (isProjectError) {
		return renderError(isProjectError);
	}

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<ul>
					{badAlbumList.map((badAlbum) => (
						<BadAlbumInfo badAlbum={badAlbum} />
					))}
				</ul>
			</section>
		</>
	);
}

// render the album information.
const BadAlbumInfo = ({ badAlbum }) => {
	return (
		<li >
			{badAlbum.artist_name}, {badAlbum.album_name}, {badAlbum.format}
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
