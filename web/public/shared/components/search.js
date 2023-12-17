// @ts-nocheck
import styles from "./search.module.css";
import classnames from "classnames";

export default (props) => (
	<input
		type="text"
		value={props.value}
		ref={props.ref}
		className={classnames(styles.root, props.className)}
		onClick={props.onClick}
		onChange={props.onChange}
		onMouseEnter={props.onMouseEnter}
		onMouseLeave={props.onMouseLeave}
		disabled={props.disabled}
		spellcheck={props.spellCheck}
		placeholder={props.placeholder}
	/>
);
