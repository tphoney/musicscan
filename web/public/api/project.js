import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * createProject creates a new project.
 */
export const createProject = async (data, fetcher) => {
	return fetcher(`${instance}/api/v1/projects`, {
		body: JSON.stringify(data),
		method: "POST",
	}).then((project) => {
		mutate(`${instance}/api/v1/user/projects`);
		return project;
	});
};

/**
 * updateProject updates an existing project
 */
export const updateProject = (params, data, fetcher) => {
	const { id } = params;
	return fetcher(`${instance}/api/v1/projects/${id}`, {
		body: JSON.stringify(data),
		method: "PATCH",
	}).then((_) => {
		mutate(`${instance}/api/v1/user/projects`);
		mutate(`${instance}/api/v1/user/projects/${id}`);
		return;
	});
};

/**
 * deleteProject deletes an existing project
 */
export const deleteProject = (params, fetcher) => {
	const { id } = params;
	return fetcher(`${instance}/api/v1/projects/${id}`, {
		method: "DELETE",
	}).then((_) => {
		mutate(`${instance}/api/v1/user/projects`);
		return;
	});
};

/**
 * useProjectList returns an swr hook that provides a project list.
 */
export const useProjectList = () => {
	const { data, error } = useSWR(`${instance}/api/v1/user/projects`);
	return {
		projectList: data,
		isLoading: !error && !data,
		isError: error,
	};
};

/**
 * useProject returns an swr hook that provides the project.
 */
export const useProject = (id) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${id}?token=true`
	);
	return {
		project: data,
		isLoading: !error && !data,
		isError: error,
	};
};

export const useAlbumBadList = (id) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${id}/bad_albums`
	);

	return {
		badAlbumList: data,
		isLoading: !error && !data,
		isError: error,
	};
};

export const useAlbumWantedList = (id, year) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${id}/wanted_albums/${year}`
	);

	return {
		wantedAlbumList: data,
		isLoading: !error && !data,
		isError: error,
	};
};

export const useRecommendedArtistList = (id) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${id}/recommended_artists`
	);

	return {
		recommendedArtistList: data,
		isLoading: !error && !data,
		isError: error,
	};
};

export const useScan = (params, scanFolder, fetcher) => {
	const { id } = params;
	return fetcher(`${instance}/api/v1/projects/${id}/scan?scan_folder=${scanFolder}`, {});
};

export const useLookup = (params, spotifyClient, spotifySecret, fetcher) => {
	const { id } = params;
	return fetcher(`${instance}/api/v1/projects/${id}/artists/lookup?spotify_client=${spotifyClient}&spotify_secret=${spotifySecret}`, {});
};