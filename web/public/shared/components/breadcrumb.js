import styles from "./breadcrumb.module.css";
import classnames from "classnames";

export default (props) => {
	const separator = props.separator || <Separator />;
	const children =
		(props.children &&
			props.children.reduce &&
			props.children.reduce((accumulator, item) => {
				accumulator.push(item);
				accumulator.push(separator);
				return accumulator;
			}, [])) ||
		props.children;
	return (
		<div className={classnames(styles.root, props.className)}>{children}</div>
	);
};

const Separator = () => (
	<svg
		width="24"
		height="24"
		viewBox="0 0 24 24"
		fill="none"
		stroke="currentColor"
		strokeWidth="2"
		strokeLinecap="round"
		strokeLinejoin="round"
		className={styles.separator}
	>
		<polyline points="9 18 15 12 9 6"></polyline>
	</svg>
);
