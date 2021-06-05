import { useState, useRef } from "react";
import styles from "./album_list.module.css";
import { Link } from "wouter";
import { useSession } from "../hooks/session.js";
import { useProject, useScan } from "../api/project";

// Renders the Album List page.
export default function ProjectScan({ params }) {
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

	const handleScan = () => {
		var data = useScan(project, fetcher);
		return (<>
			<section className={styles.root}>
				<div>Scan complete</div>
			</section>
		</>);
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
			<button onClick={handleScan}>Scan Project</button>
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
