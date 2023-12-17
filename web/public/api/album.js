import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * createAlbum creates a new album.
 */
export const createAlbum = async (project, artist, data, fetcher) => {
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
 * updateAlbum updates an existing album.
 */
export const updateAlbum = (project, artist, album, data, fetcher) => {
	return fetcher(
		`${instance}/api/v1/projects/${project}/artists/${artist}/albums/${album}`,
		{
			body: JSON.stringify(data),
			method: "PATCH",
		}
	);
};

/**
 * deleteAlbum deletes an existing album.
 */
export const deleteAlbum = (project, artist, album, fetcher) => {
	return fetcher(
		`${instance}/api/v1/projects/${project}/artists/${artist}/albums/${album}`,
		{
			method: "DELETE",
		}
	).then((_) => {
		mutate(`${instance}/api/v1/projects/${project}/artists/${artist}/albums`);
		return;
	});
};

/**
 * use returns an swr hook that provides
 */
export const useAlbumList = (project, artist) => {
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
export const useAlbum = (project, artist, album) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/artists/${artist}/albums/${album}`
	);

	return {
		album: data,
		isLoading: !error && !data,
		isError: error,
	};
};
