'use client'

import React from 'react';

const oauthMethodList = [
  { name: 'github', authUrl: 'http://localhost:3333/v1/auth/github' },
];

const OauthButtons = () => {
  return (
    <div>
      <button onClick={() => console.log('hi')}>hi</button>
      {oauthMethodList.map(({ authUrl, name }) => {
        return (
          <button
            key={name}
            onClick={() => {
              console.log('hello');
              window.open(authUrl);
            }}
          >
            {name}
          </button>
        );
      })}
    </div>
  );
};

export default OauthButtons;
