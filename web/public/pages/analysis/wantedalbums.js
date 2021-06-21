import { useState, useRef } from "react";
import styles from "./wantedalbums.module.css";
import { Link } from "wouter";

import {useAlbumWantedList, useProject } from "../../api/project.js";
import { useSession } from "../../hooks/session.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";


import { Drawer, Target } from "@accessible/drawer";

// Renders the Album List page.
export default function WantedAlbumList({ params }) {
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
		wantedAlbumList,
		isLoading: isAlbumLoading,
		isError: isAlbumError,
	} = useAlbumWantedList(params.project, params.year);

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
				<h1>Wanted Albums</h1>
				<ul className={styles.list}>
					{wantedAlbumList.map((wantedAlbum) => (
						<WantedAlbumInfo
						wantedAlbum={wantedAlbum}
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
const WantedAlbumInfo = ({ wantedAlbum }) => {
	return (
		<li id={wantedAlbum.artist_name} className={styles.item}>
			<Avatar text={wantedAlbum.artist_name} className={styles.avatar} />
			{wantedAlbum.artist_name}, {wantedAlbum.album_name}, {wantedAlbum.format}
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
