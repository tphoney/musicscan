// @ts-nocheck
import { forwardRef } from "react";
import styles from "./textarea.module.css";
import classnames from "classnames";

export default forwardRef((props, ref) => {
	return (
		<textarea
			className={classnames(styles.root, props.className)}
			ref={ref}
			onClick={props.onClick}
			onChange={props.onChange}
			onMouseEnter={props.onMouseEnter}
			onMouseLeave={props.onMouseLeave}
			disabled={props.disabled}
			spellcheck={props.spellCheck}
			placeholder={props.placeholder}
		>
			{props.children}
		</textarea>
	);
});
