import { Link as RouterLink, useRoute } from "wouter";

// Renders the Account page.
export default function Link(props) {
	const [isActive] = useRoute(props.href);
	return (
		<RouterLink {...props}>
			<a className={isActive ? "active" : ""} data-active={isActive}>
				{props.children}
			</a>
		</RouterLink>
	);
}
