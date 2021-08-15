import { useState, useRef } from "react";
import styles from "./recommendedartists.module.css";
import { Link } from "wouter";

import { useRecommendedArtistList, useProject } from "../../api/project.js";
import { useSession } from "../../hooks/session.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";


import { Drawer, Target } from "@accessible/drawer";

// Renders the Album List page.
export default function RecommendedArtistList({ params }) {
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
	// Load Artist List
	//
	const {
		recommendedArtistList,
		isLoading: isAlbumLoading,
		isError: isAlbumError,
	} = useRecommendedArtistList(params.project);

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
				<h1>Recommended Artists</h1>
				<ul className={styles.list}>
					{recommendedArtistList.map((recommendedArtist) => (
						<RecommendedArtistInfo project={project}
							recommendedArtist={recommendedArtist}
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

// render the artist information.
const RecommendedArtistInfo = ({ project, recommendedArtist }) => {
	return (
		<li id={recommendedArtist.name} className={styles.item}>
			<Avatar text={recommendedArtist.name} className={styles.avatar} />
			{recommendedArtist.name}, {recommendedArtist.popularity}
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
