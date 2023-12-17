import { useState } from "react";
import styles from "./album.module.css";
import { useSession } from "../../hooks/session.js";
import { useProject } from "../../api/project.js";
import { useArtist } from "../../api/artist.js";
import { useAlbum } from "../../api/album.js";
import { Link } from "wouter";

import Button from "../../shared/components/button";
import Breadcrumb from "../../shared/components/breadcrumb";
import Input from "../../shared/components/input";

// Renders the Album Info page.
export default function Album({ params }) {
	const { session, fetcher } = useSession();
	const [showToken, setShowToken] = useState(false);

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
		album,
		isLoading: isAlbumLoading,
		isError: isAlbumError,
	} = useAlbum(params.project, params.artist, params.album);

	if (isAlbumLoading) {
		return renderLoading();
	}
	if (isAlbumError) {
		return renderError(isAlbumError);
	}

	return (
		<>
			<section className={styles.root}>
				<Breadcrumb>
					<Link href={`/projects/${project.id}/artists`}>Artists</Link>
					<Link href={`/projects/${project.id}/artists/${artist && artist.id}`}>
						{artist && artist.name}
					</Link>
					<Link href={`/projects/${project.id}/artists/${artist && artist.id}`}>
						Albums
					</Link>
				</Breadcrumb>
				<h1>Album</h1>
				<div className={styles.card}>
					<h2>Album</h2>
					<div className={styles.field}>
						<label>Name *</label>
						<Input type="text" value={album && album.name} />
					</div>
					<div className={styles.field}>
						<label>Description *</label>
						<Input type="text" value={album && album.desc} />
					</div>
					<div className={styles.actions}>
						<Button>Update</Button>
					</div>
				</div>

				<div className={styles.card}>
					<h2>Delete Album</h2>
					<p>Warning, this action cannot be undone.</p>
					<Button>Delete</Button>
				</div>
			</section>
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
