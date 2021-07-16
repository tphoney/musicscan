import { useState, useRef } from "react";
import styles from "./albums.module.css";
import { Link } from "wouter";
import { useArtist } from "../../api/artist.js";
import {
	useAlbumList,
	createAlbum,
	deleteAlbum,
} from "../../api/album.js";
import { useProject } from "../../api/project.js";
import { useSession } from "../../hooks/session.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";
import Breadcrumb from "../../shared/components/breadcrumb";

import { Drawer, Target } from "@accessible/drawer";

// Renders the Album List page.
export default function AlbumList({ params }) {
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
		createAlbum(project.id, artist.id, { name, desc }, fetcher).then(
			(album) => {
				nameElem.current.value = "";
				descElem.current.value = "";
				setOpen(false);
			}
		);
	};

	//
	// Delete Album Functions
	//

	const handleDelete = (project, artist, album) => {
		deleteAlbum(project.id, artist.id, album.id, fetcher);
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<Breadcrumb>
					<Link href={`/projects/${project.id}/artists`}>Artists</Link>
					<Link href={`/projects/${project.id}/artists/${artist && artist.id}`}>
						{artist && artist.name}
					</Link>
				</Breadcrumb>

				<h1>Albums</h1>
				<ul className={styles.list}>
					{albumList.map((album) => (
						<AlbumInfo
							artist={artist}
							album={album}
							project={project}
							onDelete={handleDelete}
						/>
					))}
				</ul>

				<Button className={styles.button} onClick={() => setOpen(true)}>
					New Album
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
							<Button onClick={handleCreate}>Add Album</Button>
							<Button onClick={() => setOpen(false)}>Close</Button>
						</div>
					</div>
				</Target>
			</Drawer>
		</>
	);
}

// render the album information.
const AlbumInfo = ({ artist, album, project, onDelete }) => {
	return (
		<li id={album.id} className={styles.item}>
			<Avatar text={album.name} className={styles.avatar} />
			<Link
				href={`/projects/${project.id}/artists/${artist.id}/albums/${album.id}`}
				className={styles.fill}
			>
				{album.name}
			</Link>
			<label>{album.format} {album.year}</label><input type="checkbox" checked={album.wanted}></input>
			<Button onClick={onDelete.bind(this, project, artist, album)}>
				Delete
			</Button>
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
