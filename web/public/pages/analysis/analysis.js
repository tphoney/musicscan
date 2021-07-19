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

	const handleScan = () => {
		var data = useScan(project, fetcher);
		return (<>
			<section className={styles.root}>
				<div>Scan complete</div>
			</section>
		</>);
	};

	const handleLookup = () => {
		var data = useLookup(project, fetcher);
		return (<>
			<section className={styles.root}>
				<div>Lookup complete</div>
			</section>
		</>);
	};

	const { project } = useProject(params.project);

	const [year, setYear] = useState(" ");

	const handleUpdateYear = (event) => {
		setYear(event.target.value);
	};

	const [projectData, setProjectData] = useState({});
	useEffect(() => project && setProjectData(project), [project]);

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
					<h2>Scan Disk</h2>
					<p>Scans the hard disk looking for artists and albums.</p>
					<Button onClick={handleScan}>Scan Disk</Button>
				</div>
				<div className={styles.card}>
					<h2>Spotify Lookup</h2>
					<p>Scans Spotifylooking up artists and matching albums.</p>
					<Button onClick={handleLookup}>Spotify Lookup</Button>
				</div>
			</section>
		</>
	);
}
