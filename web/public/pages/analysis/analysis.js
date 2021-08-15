import { useState, useEffect } from "react";
import { useLocation } from "wouter";
import styles from "./analysis.module.css";
import { useSession } from "../../hooks/session.js";
import { useProject, useScan, useLookup } from "../../api/project.js";

import Input from "../../shared/components/input";
import Button from "../../shared/components/button";

// Renders the Project Analysis page.
export default function Analysis({ params }) {
	const { fetcher } = useSession();
	const [location, setLocation] = useLocation();
	//
	// Load Project
	//
	const { project } = useProject(params.project);

	const [scanFolder, setScanFolder] = useState("/media/tp/stuff/Music");

	const handleUpdateScanFolder = (event) => {
		setScanFolder(event.target.value);
	};

	const [spotifyKey, setSpotifyKey] = useState("");

	const handleUpdateSpotifyKey = (event) => {
		setSpotifyKey(event.target.value);
	};

	const [year, setYear] = useState("2021");

	const handleUpdateYear = (event) => {
		setYear(event.target.value);
	};

	const [projectData, setProjectData] = useState({});
	useEffect(() => project && setProjectData(project), [project]);

	const handleScan = () => {
		 if (scanFolder == "") {
			window.alert(`Please select a scan folder`);
			return;
		 }
		var data = useScan(project, scanFolder, fetcher);
		return (<>
			window.
		</>);
	};

	const handleLookup = () => {
		if (spotifyKey == "") {
			window.alert(`Please insert your spotify API key`);
			return;
		 }
		var data = useLookup(project, spotifyKey, fetcher);
		return (<>
			<section className={styles.root}>
				<div>Lookup complete</div>
			</section>
		</>);
	};

	return (
		<>
			<section className={styles.root}>
				<div className={styles.card}>
					<h2>Bad albums</h2>
					<p>MP3 albums that should be replaced.</p>
					<Button onClick={() => setLocation(`/projects/${project.id}/analysis/bad_album_list`)}>Bad Albums</Button>
				</div>
				<div className={styles.card}>
					<h2>Wanted albums</h2>
					<div className={styles.field}>
						<label>Year</label>
						<Input type="text" onChange={handleUpdateYear} placeholder="2021" />
					</div>
					<p>List unowned albums for a year</p>
					<Button onClick={() => setLocation(`/projects/${project.id}/analysis/wanted_album_list/${year}`)}>Wanted albums</Button>
				</div>
				<div className={styles.card}>
					<h2>Recommended Artists</h2>
					<p>List of artits you do not have currently.</p>
					<Button onClick={() => setLocation(`/projects/${project.id}/analysis/recommended_artist_list`)}>Recommended Artists</Button>
				</div>
				<div className={styles.card}>
					<h2>Scan Disk</h2>
					<div className={styles.field}>
						<label>Folder</label>
						<Input type="text" onChange={handleUpdateScanFolder} placeholder="/media/tp/stuff/Music" />
					</div>
					<p>Scans the hard disk looking for artists and albums.</p>
					<Button onClick={handleScan}>Scan Disk</Button>
				</div>
				<div className={styles.card}>
					<h2>Spotify Lookup</h2>
					<Input type="text" onChange={handleUpdateSpotifyKey} />
					<p>Scans Spotify looking up artists/albums and Similar Artists.</p>
					<Button onClick={handleLookup}>Spotify Lookup</Button>
				</div>
			</section>
		</>
	);
}
