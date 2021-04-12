import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * createMember creates a new member.
 */
export const createMember = (project, data, fetcher) => {
	return fetcher(
		`${instance}/api/v1/projects/${project}/members/${data.email}`,
		{
			body: JSON.stringify(data),
			method: "POST",
		}
	).then((member) => {
		mutate(`${instance}/api/v1/projects/${project}/members`);
		return member;
	});
};

/**
 * updateMember updates an existing member.
 */
export const updateMember = (project, member, data, fetcher) => {
	return fetcher(`${instance}/api/v1/projects/${project}/members/${member}`, {
		body: JSON.stringify(data),
		method: "PATCH",
	});
};

/**
 * deleteMember deletes an existing member.
 */
export const deleteMember = (project, member, fetcher) => {
	return fetcher(`${instance}/api/v1/projects/${project}/members/${member}`, {
		method: "DELETE",
	});
};

/**
 * use returns an swr hook that provides
 */
export const useMemberList = (project) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/members`
	);

	return {
		memberList: data,
		isLoading: !error && !data,
		isError: error,
	};
};
