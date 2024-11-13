interface UserListProps {
    online: boolean;
}

export const UserOnlineIcon = ({ online }: UserListProps) => {
    return (
        <svg
            width="13"
            height="12"
            viewBox="0 0 13 12"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
        >
            <ellipse
                id="Ellipse 3"
                cx="6.45222"
                cy="5.95117"
                rx="5.57478"
                ry="5.78125"
                fill={online ? "#17C900" : "#FF0000"}
            />{" "}
            {/*"#17C900" */}
        </svg>
    );
};
