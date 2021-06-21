import { useState, useRef } from "react";
import styles from "./artists.module.css";
import { Link } from "wouter";
import { useArtistList, createArtist, deleteArtist } from "../../api/artist.js";
import { useProject } from "../../api/project.js";
import { useSession } from "../../hooks/session.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";

import { Drawer, Target } from "@accessible/drawer";

// Renders the Artist List page.
export default function ArtistList({ params }) {
	const { fetcher } = useSession();
	const [open, setOpen] = useState(false);

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
			setOpen(false);
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
				<ul className={styles.list}>
					{artistList.map((artist) => (
						<ArtistInfo
							artist={artist}
							project={project}
							onDelete={handleDelete}
						/>
					))}
				</ul>

				<Button className={styles.button} onClick={() => setOpen(true)}>
					New Artist
				</Button>
			</section>

			<Drawer open={open}>
				<Target
					placement="right"
					closeOnEscape={true}
					preventScroll={true}
					openClass={styles.drawer}
				>
					<div>
						<Input ref={nameElem} type="text" placeholder="name" />
						<Input ref={descElem} type="text" placeholder="desc" />

						<div className={styles.actions}>
							<Button onClick={handleCreate}>Add Artist</Button>
							<Button onClick={() => setOpen(false)}>Close</Button>
						</div>
					</div>
				</Target>
			</Drawer>
		</>
	);
}

// render the artist information.
const ArtistInfo = ({ artist, project, onDelete }) => {
	return (
		<li id={artist.id} className={styles.item}>
			<Avatar text={artist.name} className={styles.avatar} />
			<Link
				href={`/projects/${project.id}/artists/${artist.id}`}
				className={styles.fill}
			>
				{artist.name}
			</Link>
			<input type="checkbox" checked={artist.wanted}></input>
			<Button onClick={onDelete.bind(this, artist)}>Delete</Button>
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
