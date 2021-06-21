import styles from "./checkbox.module.css";
import classnames from "classnames";

import * as Checkbox from "@radix-ui/react-checkbox";

export default (props) => (
	<Checkbox.Root
		className={classnames(styles.root, props.className)}
		onCheckedChange={props.onCheckedChange}
		disabled={props.disabled}
		checked={props.checked}
	>
		<Checkbox.Indicator as={Checkmark} />
	</Checkbox.Root>
);

const Checkmark = () => (
	<svg
		width="24"
		height="24"
		viewBox="0 0 24 24"
		fill="none"
		stroke="currentColor"
		strokeWidth="4"
		strokeLinecap="round"
		strokeLinejoin="round"
	>
		<polyline points="20 6 9 17 4 12"></polyline>
	</svg>
);
