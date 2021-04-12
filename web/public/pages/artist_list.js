import { useState, useRef } from "react";
import styles from "./artist_list.module.css";
import { Link } from "wouter";
import { useartistList, createartist, deleteartist } from "../api/artist.js";
import { useProject } from "../api/project.js";
import { useSession } from "../hooks/session.js";

// Renders the artist List page.
export default function artistList({ params }) {
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
	// Load artist List
	//

	const {
		artistList,
		isLoading: isartistLoading,
		isError: isartistErrror,
	} = useartistList(project && project.id);

	if (isartistLoading) {
		return renderLoading();
	}
	if (isartistErrror) {
		return renderError(isartistErrror);
	}

	//
	// Add artist Functions
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		const data = { name, desc };
		const params = { project: project.id };
		createartist(params, data, fetcher).then((project) => {
			nameElem.current.value = "";
			descElem.current.value = "";
		});
	};

	//
	// Handle Deletions
	//

	const handleDelete = (artist) => {
		const params = { project: project.id, artist: artist.id };
		deleteartist(params, fetcher);
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<ul>
					{artistList.map((artist) => (
						<artistInfo
							artist={artist}
							project={project}
							onDelete={handleDelete}
						/>
					))}
				</ul>

				<div className="actions">
					<button onClick={handleCreate}>Add artist</button>
					<input ref={nameElem} type="text" placeholder="name" />
					<input ref={descElem} type="text" placeholder="desc" />
				</div>
			</section>
		</>
	);
}

// render the artist information.
const artistInfo = ({ artist, project, onDelete }) => {
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
	return <div>Your artist list is empty</div>;
};
