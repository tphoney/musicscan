import { useState, useRef } from "react";
import styles from "./album_list.module.css";
import { Link } from "wouter";
import { useArtist } from "../api/artist.js";
import { useAlbumList, createAlbum } from "../api/album.js";
import { useProject } from "../api/project.js";
import { useSession } from "../hooks/session.js";

// Renders the Album List page.
export default function AlbumList({ params }) {
	const { fetcher } = useSession();

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
	// Load Album List
	//

	const {
		albumList,
		isLoading: isAlbumLoading,
		isError: isAlbumError,
	} = useAlbumList(params.project, params.artist);

	if (isAlbumLoading) {
		return renderLoading();
	}
	if (isAlbumError) {
		return renderError(isAlbumError);
	}

	//
	// Add Album Functions
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		createAlbum(project.id, artist.id, { name, desc }, fetcher).then((album) => {
			nameElem.current.value = "";
			descElem.current.value = "";
		});
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<ul>
					{albumList.map((album) => (
						<AlbumInfo artist={artist} album={album} project={project} />
					))}
				</ul>

				<div className="actions">
					<button onClick={handleCreate}>Add Album</button>
					<input ref={nameElem} type="text" placeholder="name" />
					<input ref={descElem} type="text" placeholder="desc" />
				</div>
			</section>
		</>
	);
}

// render the album information.
const AlbumInfo = ({ artist, album, project }) => {
	return (
		<li id={artist.id}>
			<Link
				href={`/projects/${project.id}/artists/${artist.id}/albums/${album.id}`}
			>
				{album.name}
			</Link>
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
