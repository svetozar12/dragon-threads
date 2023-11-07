import React from 'react';
import styles from './page.module.css';
import Button from '@mui/material/Button';
export default async function Index() {
  /*
   * Replace the elements below with your own.
   *
   * Note: The corresponding styles are in the ./index.css file.
   */
  return (
    <div className={styles.page}>
      <Button variant="contained">Hello world</Button>
    </div>
  );
}
