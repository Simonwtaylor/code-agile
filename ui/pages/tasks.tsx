import React from 'react';
 
export interface ITasksPageProps {
  tasks: any[];
}
 
const tasksPage: React.FC<ITasksPageProps> = ({
  tasks,
}) => {
  console.log(tasks);
  if (!tasks) {
    return <h1>Unable to fetch tasks :(</h1>
  }

  if (tasks.length === 0) {
    return <h1>Uhhh...No Tasks found</h1>
  }

  return (
    <>
      {tasks.map(({ id, title }) => <div key={id}>{title}</div>)}
    </>
  );
};
 
export default tasksPage;

export async function getServerSideProps(context: any) {
  const res = await fetch('http://localhost:3001/tasks');

  if (res.status >= 400) {
    console.log('Issue fetching tasks');
  }

  const data = await res.json();

  return {
    props: {
      tasks: data,
    },
  }
}