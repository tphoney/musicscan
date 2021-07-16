import { useState, useEffect } from "react";
import { useLocation } from "wouter";
import styles from "./settings.module.css";
import { useSession } from "../../hooks/session.js";
import { useProject } from "../../api/project.js";
import { useArtist, updateArtist, deleteArtist } from "../../api/artist.js";

import Button from "../../shared/components/button";
import Input from "../../shared/components/input";

// Renders the Project Settings page.
export default function Settings({ params }) {
	const { fetcher } = useSession();
	const [_, setLocation] = useLocation();

	//
	// Load Project
	//

	const { project } = useProject(params.project);

	//
	// Load Artist
	//

	const { artist } = useArtist(params.project, params.artist);
	const [artistData, setArtistData] = useState({});
	useEffect(() => artist && setArtistData(artist), [artist]);

	//
	// Update Artist
	//

	const handleUpdateName = (event) => {
		setArtistData({
			name: event.target.value,
			desc: artistData.desc,
		});
	};

	const handleUpdateDesc = (event) => {
		setArtistData({
			desc: event.target.value,
			name: artistData.name,
		});
	};

	const handleUpdate = () => {
		const params = {
			project: project.id,
			artist: artist.id,
		};
		updateArtist(params, artistData, fetcher);
	};

	//
	// Delete Artist
	//

	const handleDelete = () => {
		if (confirm("Are you sure you want to proceed?")) {
			const params = {
				project: project.id,
				artist: artist.id,
			};
			deleteArtist(params, fetcher);
			setLocation(`/projects/${project.id}`);
		}
	};

	return (
		<>
			<section className={styles.root}>
				<div className={styles.card}>
					<h2>Artist</h2>
					<div className={styles.field}>
						<label>Name *</label>
						<Input
							type="text"
							value={artistData.name}
							onChange={handleUpdateName}
						/>
					</div>
					<div className={styles.field}>
						<label>Description *</label>
						<Input
							type="text"
							value={artistData.desc}
							onChange={handleUpdateDesc}
						/>
					</div>
					<div className={styles.actions}>
						<Button onClick={handleUpdate}>Update Project</Button>
					</div>
				</div>

				<div className={styles.card}>
					<h2>Delete</h2>
					<p>Warning, this action cannot be undone.</p>
					<Button onClick={handleDelete}>Delete</Button>
				</div>
			</section>
		</>
	);
}
