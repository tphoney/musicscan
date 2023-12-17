import { useState, useRef } from "react";
import styles from "./projects.module.css";
import { Link } from "wouter";
import {
	useProjectList,
	createProject,
	deleteProject,
} from "../../api/project.js";
import { useSession } from "../../hooks/session.js";

import Button from "../../shared/components/button";
import Input from "../../shared/components/input";
import Avatar from "../../shared/components/avatar";

import { Drawer, Target } from "@accessible/drawer";

// Renders the Home page.
export default function Home() {
	const { fetcher } = useSession();
	const [open, setOpen] = useState(false);

	//
	// Load Project List
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
			setOpen(false);
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
				<ul className={styles.list}>
					{projectList.map((project) => (
						<ProjectInfo project={project} onDelete={handleDelete} />
					))}
				</ul>

				<Button className={styles.button} onClick={() => setOpen(true)}>
					New Project
				</Button>

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
								<Button onClick={handleCreate}>Create Project</Button>
								<Button onClick={() => setOpen(false)}>Close</Button>
							</div>
						</div>
					</Target>
				</Drawer>
			</section>
		</>
	);
}

// render the project information.
const ProjectInfo = ({ project, onDelete }) => {
	return (
		// <li >
		<Link
			href={`/projects/${project.id}`}
			id={project.id}
			className={styles.item}
		>
			<Avatar text={project.name} />
			{project.name}
		</Link>
		///* <button onClick={onDelete.bind(this, project)}>Delete</button> */}
		///* </li> */}
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
