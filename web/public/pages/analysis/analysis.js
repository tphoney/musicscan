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

	const [spotifyClient, setSpotifyClient] = useState("4bfbcdb641bb4409ad6e1a39a67d752b");

	const handleUpdateSpotifyClient = (event) => {
		setSpotifyClient(event.target.value);
	};

	const [spotifySecret, setSpotifySecret] = useState("d5239a128ef443a68d6fef2aa609920e");

	const handleUpdateSpotifySecret = (event) => {
		setSpotifySecret(event.target.value);
	};

	// use the current year as the default
	const [year, setYear] = useState(new Date().getFullYear());

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
		if (spotifyClient == "") {
			window.alert(`Please insert your spotify client`);
			return;
		}
		if (spotifySecret == "") {
			window.alert(`Please insert your spotify secret`);
			return;
		}
		var data = useLookup(project, spotifyClient, spotifySecret, fetcher);
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
						<Input type="text" onChange={handleUpdateYear} placeholder="2023" />
					</div>
					<p>List unowned albums for a year</p>
					<Button onClick={() => setLocation(`/projects/${project.id}/analysis/wanted_album_list/${year}`)}>Wanted albums</Button>
				</div>
				<div className={styles.card}>
					<h2>Recommended Artists</h2>
					<p>List of artists you do not have currently.</p>
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
					<div className={styles.field}>
						<label>client id</label>
						<Input type="text" onChange={handleUpdateSpotifyClient} />
					</div>
					<div className={styles.field}>
						<label>secret</label>	
						<Input type="text" onChange={handleUpdateSpotifySecret} />
					</div>
					<p>Scans Spotify looking up artists/albums and Similar Artists.</p>
					<Button onClick={handleLookup}>Spotify Lookup</Button>
				</div>
			</section>
		</>
	);
}
