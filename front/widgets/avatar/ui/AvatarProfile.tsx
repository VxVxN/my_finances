import { LoginButton } from "@/features/LoginButton";
import {
  Avatar,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownTrigger,
} from "@nextui-org/react";

export const AvatarProfile = () => {
    const isAuth = false

    if (!isAuth) {
        return (<LoginButton/>)
    }

  return (
    <Dropdown placement="bottom-end">
      <DropdownTrigger>
        <Avatar
          isBordered
          as="button"
          className="transition-transform"
          color="secondary"
          name="Jason Hughes"
          size="sm"
          src="https://i.pravatar.cc/150?u=a042581f4e29026704d"
        />
      </DropdownTrigger>
      <DropdownMenu aria-label="Profile Actions" variant="flat"> 
        <DropdownItem key="settings">Настройки</DropdownItem> 
        <DropdownItem key="logout" color="danger">
          Выход
        </DropdownItem>
      </DropdownMenu>
    </Dropdown>
  );
};
