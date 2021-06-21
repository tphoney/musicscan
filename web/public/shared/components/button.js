import styles from "./button.module.css";
import classnames from "classnames";

export default (props) => (
	<button
		className={classnames(styles.root, props.className)}
		onClick={props.onClick}
		onMouseEnter={props.onMouseEnter}
		onMouseLeave={props.onMouseLeave}
		disabled={props.disabled}
	>
		{props.children}
	</button>
);
