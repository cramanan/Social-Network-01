import React from "react";

interface ProfileCircleProps {
  className?: string;
}

export const ProfileCircle: React.FC<ProfileCircleProps> = ({ className }) => {
  return (
    <div className={`profile-circle ${className}`}>
      <svg
        width="44"
        height="44"
        viewBox="0 0 62 62"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <g id="vuesax/linear/profile-circle">
          <g id="profile-circle">
            <path
              id="Vector"
              d="M31.31 33.0151C31.1292 32.9892 30.8967 32.9892 30.69 33.0151C26.1433 32.8601 22.5267 29.1401 22.5267 24.5676C22.5267 19.8917 26.2983 16.0942 31 16.0942C35.6758 16.0942 39.4733 19.8917 39.4733 24.5676C39.4475 29.1401 35.8567 32.8601 31.31 33.0151Z"
              stroke="url(#paint0_linear_380_404)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              id="Vector_2"
              d="M48.4117 50.0648C43.8133 54.2756 37.7167 56.8331 31 56.8331C24.2833 56.8331 18.1867 54.2756 13.5883 50.0648C13.8467 47.6365 15.3967 45.2598 18.1608 43.3998C25.2392 38.6981 36.8125 38.6981 43.8392 43.3998C46.6033 45.2598 48.1533 47.6365 48.4117 50.0648Z"
              stroke="url(#paint1_linear_380_404)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              id="Vector_3"
              d="M31 56.8332C45.2674 56.8332 56.8334 45.2672 56.8334 30.9998C56.8334 16.7325 45.2674 5.1665 31 5.1665C16.7327 5.1665 5.16669 16.7325 5.16669 30.9998C5.16669 45.2672 16.7327 56.8332 31 56.8332Z"
              stroke="url(#paint2_linear_380_404)"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </g>
        </g>
        <defs>
          <linearGradient
            id="paint0_linear_380_404"
            x1="22.5267"
            y1="16.4069"
            x2="37.4546"
            y2="3.16061"
            gradientUnits="userSpaceOnUse"
          >
            <stop offset="0.00899061" stopColor="#4921FA" stopOpacity="0.4" />
            <stop offset="0.195904" stopColor="#6F47C0" stopOpacity="0.42" />
            <stop offset="0.340981" stopColor="#B269E2" stopOpacity="0.53" />
            <stop offset="0.6244" stopColor="#E1D3EB" stopOpacity="0.67" />
          </linearGradient>
          <linearGradient
            id="paint1_linear_380_404"
            x1="13.5883"
            y1="40.187"
            x2="26.3107"
            y2="17.042"
            gradientUnits="userSpaceOnUse"
          >
            <stop offset="0.00899061" stopColor="#4921FA" stopOpacity="0.4" />
            <stop offset="0.195904" stopColor="#6F47C0" stopOpacity="0.42" />
            <stop offset="0.340981" stopColor="#B269E2" stopOpacity="0.53" />
            <stop offset="0.6244" stopColor="#E1D3EB" stopOpacity="0.67" />
          </linearGradient>
          <linearGradient
            id="paint2_linear_380_404"
            x1="5.16669"
            y1="6.12131"
            x2="50.7399"
            y2="-34.2565"
            gradientUnits="userSpaceOnUse"
          >
            <stop offset="0.00899061" stopColor="#4921FA" stopOpacity="0.4" />
            <stop offset="0.195904" stopColor="#6F47C0" stopOpacity="0.42" />
            <stop offset="0.340981" stopColor="#B269E2" stopOpacity="0.53" />
            <stop offset="0.6244" stopColor="#E1D3EB" stopOpacity="0.67" />
          </linearGradient>
        </defs>
      </svg>
    </div>
  );
};
