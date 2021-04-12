import { useState, useRef } from "react";
import styles from "./home.module.css";
import { Link } from "wouter";
import {
	useProjectList,
	createProject,
	deleteProject,
} from "../api/project.js";
import { useSession } from "../hooks/session.js";

// Renders the Home page.
export default function Home() {
	const { fetcher } = useSession();

	//
	// Load User List
	//

	const { projectList, isLoading, isError } = useProjectList();
	if (isLoading) {
		return renderLoading();
	}
	if (isError) {
		return renderError(isError);
	}

	//
	// Create Project Function
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		createProject({ name, desc }, fetcher).then((project) => {
			nameElem.current.value = "";
			descElem.current.value = "";
		});
	};

	//
	// Handle Deletions
	//

	const handleDelete = (project) => {
		deleteProject(project, fetcher);
	};

	return (
		<>
			<section className={styles.root}>
				<h2>Projects</h2>
				<ul>
					{projectList.map((project) => (
						<ProjectInfo project={project} onDelete={handleDelete} />
					))}
				</ul>

				<div className="actions">
					<button onClick={handleCreate}>Add Project</button>
					<input ref={nameElem} type="text" placeholder="name" />
					<input ref={descElem} type="text" placeholder="desc" />
				</div>
			</section>
		</>
	);
}

// render the project information.
const ProjectInfo = ({ project, onDelete }) => {
	return (
		<li id={project.id}>
			<Link href={`/projects/${project.id}`}>{project.name}</Link>
			<button onClick={onDelete.bind(this, project)}>Delete</button>
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
	return <div>Your Project list is empty</div>;
};
