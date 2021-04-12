import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * createalbum creates a new album.
 */
export const createalbum = async (project, artist, data, fetcher) => {
	return fetcher(
		`${instance}/api/v1/projects/${project}/artists/${artist}/albums`,
		{
			body: JSON.stringify(data),
			method: "POST",
		}
	).then((album) => {
		mutate(`${instance}/api/v1/projects/${project}/artists/${artist}/albums`);
		return album;
	});
};

/**
 * updatealbum updates an existing album.
 */
export const updatealbum = (project, artist, album, data, fetcher) => {
	return fetcher(
		`${instance}/api/v1/projects/${project}/artists/${artist}/albums/${album}`,
		{
			body: JSON.stringify(data),
			method: "PATCH",
		}
	);
};

/**
 * deletealbum deletes an existing album.
 */
export const deletealbum = (project, artist, album, fetcher) => {
	return fetcher(
		`${instance}/api/v1/projects/${project}/artists/${artist}/albums/${album}`,
		{
			method: "DELETE",
		}
	);
};

/**
 * use returns an swr hook that provides
 */
export const usealbumList = (project, artist) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/artists/${artist}/albums`
	);

	return {
		albumList: data,
		isLoading: !error && !data,
		isError: error,
	};
};

/**
 * use returns an swr hook that provides
 */
export const usealbum = (project, artist, album) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/artists/${artist}/albums/${album}`
	);

	return {
		album: data,
		isLoading: !error && !data,
		isError: error,
	};
};
