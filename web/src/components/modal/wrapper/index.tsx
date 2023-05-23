import { Modal as MantineModal, useMantineTheme } from "@mantine/core";

export type ModalContainerProps = {
  opened: boolean;
  close: () => void;
  children: React.ReactNode;
};

const ModalContainer = ({ opened, close, children }: ModalContainerProps) => {
  const theme = useMantineTheme();
  return (
    <MantineModal
      opened={opened}
      onClose={close}
      title="Authentication"
      overlayProps={{
        color:
          theme.colorScheme === "dark"
            ? theme.colors.dark[9]
            : theme.colors.gray[2],
        opacity: 0.55,
        blur: 3,
      }}
    >
      {children}
    </MantineModal>
  );
};

export default ModalContainer;

// interface ModalHOCProps {
//   opened: boolean;
//   close: () => void;
//   Component: ComponentType<any>;
//   useLogic: () => {};
// }

// const ModalHOC = (Component, opened, close, useLogic) => {
//   return () => {
//     const logic = useLogic();
//     return (
//       <ModalContainer opened={opened} close={close}>
//         {/* <Component {...logic} /> */}
//         <div>hello</div>
//       </ModalContainer>
//     );
//     return <div>hello</div>;
//   };
// };

// export default ModalHOC;
