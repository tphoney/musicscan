import { useState, useRef } from "react";
import styles from "./artist_list.module.css";
import { Link } from "wouter";
import { useArtistList, createArtist, deleteArtist } from "../api/artist.js";
import { useProject } from "../api/project.js";
import { useSession } from "../hooks/session.js";

// Renders the Artist List page.
export default function ArtistList({ params }) {
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
	// Load Artist List
	//

	const {
		artistList,
		isLoading: isArtistLoading,
		isError: isArtistErrror,
	} = useArtistList(project && project.id);

	if (isArtistLoading) {
		return renderLoading();
	}
	if (isArtistErrror) {
		return renderError(isArtistErrror);
	}

	//
	// Add Artist Functions
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		const data = { name, desc };
		const params = { project: project.id };
		createArtist(params, data, fetcher).then((project) => {
			nameElem.current.value = "";
			descElem.current.value = "";
		});
	};

	//
	// Handle Deletions
	//

	const handleDelete = (artist) => {
		const params = { project: project.id, artist: artist.id };
		deleteArtist(params, fetcher);
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<ul>
					{artistList.map((artist) => (
						<ArtistInfo
							artist={artist}
							project={project}
							onDelete={handleDelete}
						/>
					))}
				</ul>

				<div className="actions">
					<button onClick={handleCreate}>Add Artist</button>
					<input ref={nameElem} type="text" placeholder="name" />
					<input ref={descElem} type="text" placeholder="desc" />
				</div>
			</section>
		</>
	);
}

// render the artist information.
const ArtistInfo = ({ artist, project, onDelete }) => {
	return (
		<li id={artist.id}>
			<Link href={`/projects/${project.id}/artists/${artist.id}`}>
				{artist.name}
			</Link>
			<button onClick={onDelete.bind(this, artist)}>Delete</button>
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

// helper function returns the empty message.
const renderEmpty = (error) => {
	return <div>Your Artist list is empty</div>;
};
