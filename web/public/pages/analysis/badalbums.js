import { useState, useRef } from "react";
import styles from "./badalbums.module.css";
import { Link } from "wouter";

import {useAlbumBadList, useProject } from "../../api/project.js";
import { useSession } from "../../hooks/session.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";


import { Drawer, Target } from "@accessible/drawer";

// Renders the Album List page.
export default function BadAlbumList({ params }) {
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
	// Load Album List
	//

	const {
		badAlbumList,
		isLoading: isAlbumLoading,
		isError: isAlbumError,
	} = useAlbumBadList(params.project);

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


	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<h1>Bad Albums</h1>
				<ul className={styles.list}>
					{badAlbumList.map((badAlbum) => (
						<BadAlbumInfo
						badAlbum={badAlbum}
						/>
					))}
				</ul>

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

							<Button onClick={() => setOpen(false)}>Close</Button>
						</div>
					</div>
				</Target>
			</Drawer>
		</>
	);
}

// render the album information.
const BadAlbumInfo = ({ badAlbum }) => {
	return (
		<li id={badAlbum.artist_name} className={styles.item}>
			<Avatar text={badAlbum.artist_name} className={styles.avatar} />
			{badAlbum.artist_name}, {badAlbum.album_name}, {badAlbum.format}
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
