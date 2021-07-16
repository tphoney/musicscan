// @ts-nocheck
import styles from "./select.module.css";
import classnames from "classnames";

export default (props) => (
	<select
		className={classnames(styles.root, props.className)}
		ref={props.ref}
		onClick={props.onClick}
		onChange={props.onChange}
		onMouseEnter={props.onMouseEnter}
		onMouseLeave={props.onMouseLeave}
		disabled={props.disabled}
		width={props.width || "350px"}
	>
		{props.children}
	</select>
);
