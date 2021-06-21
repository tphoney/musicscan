// @ts-nocheck
import { forwardRef } from "react";
import styles from "./input.module.css";
import classnames from "classnames";

export default forwardRef((props, ref) => (
	<input
		type={props.type || "text"}
		ref={ref}
		value={props.value}
		name={props.name}
		className={classnames(styles.root, props.className)}
		onClick={props.onClick}
		onChange={props.onChange}
		onMouseEnter={props.onMouseEnter}
		onMouseLeave={props.onMouseLeave}
		disabled={props.disabled}
		spellcheck={props.spellCheck}
		placeholder={props.placeholder}
	/>
));
