import { useState, useEffect } from "react";
import { useLocation } from "wouter";
import styles from "./settings.module.css";
import { useSession } from "../../hooks/session.js";
import { useProject, updateProject, deleteProject } from "../../api/project.js";

import Button from "../../shared/components/button";
import Input from "../../shared/components/input";

// Renders the Project Settings page.
export default function Settings({ params }) {
	const { fetcher } = useSession();
	const [showToken, setShowToken] = useState(false);
	const [_, setLocation] = useLocation();

	//
	// Load Project
	//

	const { project } = useProject(params.project);
	const [projectData, setProjectData] = useState({});
	useEffect(() => project && setProjectData(project), [project]);

	//
	// Update Project
	//

	const handleUpdateName = (event) => {
		setProjectData({
			name: event.target.value,
			desc: projectData.desc,
		});
	};

	const handleUpdateDesc = (event) => {
		setProjectData({
			desc: event.target.value,
			name: projectData.name,
		});
	};

	const handleUpdate = () => {
		updateProject(project, projectData, fetcher);
	};

	//
	// Delete Project
	//

	const handleDelete = () => {
		if (confirm("Are you sure you want to proceed?")) {
			deleteProject(project, fetcher);
			setLocation("/");
		}
	};

	return (
		<>
			<section className={styles.root}>
				<div className={styles.card}>
					<h2>Project</h2>
					<div className={styles.field}>
						<label>Name *</label>
						<Input
							type="text"
							value={projectData.name}
							onChange={handleUpdateName}
						/>
					</div>
					<div className={styles.field}>
						<label>Description *</label>
						<Input
							type="text"
							value={projectData.desc}
							onChange={handleUpdateDesc}
						/>
					</div>
					<div className={styles.actions}>
						<Button onClick={handleUpdate}>Update Project</Button>
					</div>
				</div>

				<div className={styles.card}>
					<h2>Token</h2>
					<p>Project access token that can be used to access the API.</p>
					{showToken && <pre>{project && project.token}</pre>}
					{!showToken && (
						<Button onClick={() => setShowToken(true)}>Display Token</Button>
					)}
				</div>

				<div className={styles.card}>
					<h2>Delete Project</h2>
					<p>Warning, this action cannot be undone.</p>
					<Button onClick={handleDelete}>Delete</Button>
				</div>
			</section>
		</>
	);
}
