import { useState, useRef } from "react";
import styles from "./members.module.css";
import { useMemberList, createMember } from "../api/member.js";
import { useProject } from "../api/project.js";
import { useSession } from "../hooks/session.js";

// Renders the Member List page.
export default function Project({ params }) {
	const { fetcher } = useSession();

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
	// Load Member List
	//

	const {
		memberList,
		isLoading: isMemberLoading,
		isError: isMemberErrror,
	} = useMemberList(project && project.id);

	if (isMemberLoading) {
		return renderLoading();
	}
	if (isMemberErrror) {
		return renderError(isMemberErrror);
	}
	if (memberList.length === 0) {
		return renderEmpty();
	}

	//
	// Add Member Functions
	//

	const [error, setError] = useState(null);
	const emailElem = useRef(null);
	const roleElem = useRef(null);

	const handleCreate = () => {
		const email = emailElem.current.value;
		const role = roleElem.current.value;
		createMember(project.id, { email, role }, fetcher).then((member) => {
			emailElem.current.value = "";
		});
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<ul>
					{memberList.map((member) => (
						<MemberInfo member={member} project={project} />
					))}
				</ul>

				<div className="actions">
					<button onClick={handleCreate}>Add Member</button>
					<input ref={emailElem} type="text" placeholder="email" />
					<select ref={roleElem}>
						<option value="developer" selected>
							developer
						</option>
						<option value="admin">admin</option>
					</select>
				</div>
			</section>
		</>
	);
}

// render the member information.
const MemberInfo = ({ member, project }) => {
	return (
		<li id={member.id}>
			<span>{member.email}</span>
			{", "}
			<span>{member.role}</span>
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

// helper function returns the empty message.
const renderEmpty = (error) => {
	return <div>Your Member list is empty</div>;
};
