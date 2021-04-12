import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * createartist creates a new artist.
 */
export const createartist = async (params, data, fetcher) => {
	const { project } = params;
	return fetcher(`${instance}/api/v1/projects/${project}/artists`, {
		body: JSON.stringify(data),
		method: "POST",
	}).then((artist) => {
		mutate(`${instance}/api/v1/projects/${project}/artists`);
		return artist;
	});
};

/**
 * updateartist updates an existing artist.
 */
export const updateartist = (params, data, fetcher) => {
	const { project, artist } = params;
	return fetcher(`${instance}/api/v1/projects/${project}/artists/${artist}`, {
		body: JSON.stringify(data),
		method: "PATCH",
	});
};

/**
 * deleteartist deletes an existing artist.
 */
export const deleteartist = (params, fetcher) => {
	const { project, artist } = params;
	return fetcher(`${instance}/api/v1/projects/${project}/artists/${artist}`, {
		method: "DELETE",
	}).then((_) => {
		mutate(`${instance}/api/v1/projects/${project}/artists`);
		return;
	});
};

/**
 * use returns an swr hook that provides
 */
export const useartistList = (project) => {
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
export const useartist = (project, artist) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/artists/${artist}`
	);

	return {
		artist: data,
		isLoading: !error && !data,
		isError: error,
	};
};
