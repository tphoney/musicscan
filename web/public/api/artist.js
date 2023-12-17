import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * createArtist creates a new artist.
 */
export const createArtist = async (params, data, fetcher) => {
	const { project } = params;
	return fetcher(`${instance}/api/v1/projects/${project}/artists`, {
		body: JSON.stringify(data),
		method: "POST",
	}).then((response) => {
		mutate(`${instance}/api/v1/projects/${project}/artists`);
		return response;
	});
};

/**
 * updateArtist updates an existing artist.
 */
export const updateArtist = (params, data, fetcher) => {
	const { project, artist } = params;
	return fetcher(`${instance}/api/v1/projects/${project}/artists/${artist}`, {
		body: JSON.stringify(data),
		method: "PATCH",
	}).then((response) => {
		mutate(`${instance}/api/v1/projects/${project}/artists`);
		mutate(`${instance}/api/v1/projects/${project}/artists/${artist}`);
		return response;
	});
};

/**
 * deleteArtist deletes an existing artist.
 */
export const deleteArtist = (params, fetcher) => {
	const { project, artist } = params;
	return fetcher(`${instance}/api/v1/projects/${project}/artists/${artist}`, {
		method: "DELETE",
	}).then((response) => {
		mutate(`${instance}/api/v1/projects/${project}/artists`);
		return response;
	});
};

/**
 * use returns an swr hook that provides
 */
export const useArtistList = (project) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/artists`
	);

	return {
		artistList: data,
		isLoading: !error && !data,
		isError: error,
	};
};

/**
 * use returns an swr hook that provides
 */
export const useArtist = (project, artist) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/artists/${artist}`
	);

	return {
		artist: data,
		isLoading: !error && !data,
		isError: error,
	};
};
