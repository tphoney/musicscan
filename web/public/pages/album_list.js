import { useState, useRef } from "react";
import styles from "./album_list.module.css";
import { Link } from "wouter";
import { useartist } from "../api/artist.js";
import { usealbumList, createalbum } from "../api/album.js";
import { useProject } from "../api/project.js";
import { useSession } from "../hooks/session.js";

// Renders the album List page.
export default function albumList({ params }) {
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
	// Load album List
	//

	const {
		albumList,
		isLoading: isalbumLoading,
		isError: isalbumError,
	} = usealbumList(params.project, params.artist);

	if (isalbumLoading) {
		return renderLoading();
	}
	if (isalbumError) {
		return renderError(isalbumError);
	}

	//
	// Add album Functions
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		createalbum(project.id, artist.id, { name, desc }, fetcher).then((album) => {
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
						<albumInfo artist={artist} album={album} project={project} />
					))}
				</ul>

				<div className="actions">
					<button onClick={handleCreate}>Add album</button>
					<input ref={nameElem} type="text" placeholder="name" />
					<input ref={descElem} type="text" placeholder="desc" />
				</div>
			</section>
		</>
	);
}

// render the album information.
const albumInfo = ({ artist, album, project }) => {
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
