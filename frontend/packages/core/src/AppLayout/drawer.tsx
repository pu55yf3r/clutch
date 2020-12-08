import * as React from "react";
import { Link as RouterLink } from "react-router-dom";
import styled from "@emotion/styled";
import {
  ClickAwayListener,
  Collapse,
  Drawer as MuiDrawer,
  List,
  ListItem,
  ListItemText,
  Paper as MuiPaper,
  Popper,
  Typography,
} from "@material-ui/core";

import { useAppContext } from "../Contexts";

import { routesByGrouping, sortedGroupings } from "./utils";

const mobileDrawerWidth = "90%";
const drawerWidth = "100px";

const DrawerPanel = styled(MuiDrawer)({
  flexShrink: 0,
  minWidth: `${drawerWidth}`,
  "@media screen and (max-width: 800px)": {
    minWidth: `${mobileDrawerWidth}`,
  },
  "div[class*='MuiDrawer-paper']": {
    minWidth: `${drawerWidth}`,
    paddingTop: "32px",
    top: "64px",
    backgroundColor: "linear-gradient(0deg, #FFFFFF, #FFFFFF), #FFFFFF",
    boxShadow: "0px 5px 15px rgba(53, 72, 212, 0.2)",
    "@media screen and (max-width: 800px)": {
      minWidth: `${mobileDrawerWidth}`,
    },
  },
});

const GroupList = styled(List)({
  paddingTop: "0px",
  paddingBottom: "0px",
});

const GroupListItem = styled(ListItem)({
  height: "82px",
  "&:hover": {
    backgroundColor: "#E7E7EA",
  },
  "&:active": {
    backgroundColor:
      "linear-gradient(0deg, rgba(255, 255, 255, 0.85), rgba(255, 255, 255, 0.85)), #0D1030",
  },
  "&.Mui-selected": {
    backgroundColor: "#EBEDFB",
  },
});

const GroupHeading = styled(Typography)({
  color: "rgba(13, 16, 48, 0.6)",
  fontWeight: 500,
  fontSize: "14px",
  lineHeight: "18px",
  flexGrow: 1,
});

const LinkListItem = styled(ListItem)((props: { isSelectedWorkflow: boolean }) => ({
  backgroundColor: props.isSelectedWorkflow ? "#EBEDFB" : "#FFFFFF",
  height: "48px",
}));

const LinkHeading = styled(Typography)((props: { isSelectedWorkflow: boolean }) => ({
  color: props.isSelectedWorkflow ? "#3548D4" : "rgba(13, 16, 48, 0.6)",
  fontWeight: 500,
  fontSize: "14px",
  lineHeight: "18px",
}));

const Paper = styled(MuiPaper)({
  minWidth: "230px",
  border: "1px solid #E7E7EA",
  boxShadow: "0px 10px 24px rgba(35, 48, 143, 0.3)",
});

interface GroupProps {
  heading: string;
  open: boolean;
  updateOpenGroup: (heading: string) => void;
  onNavigate: () => void;
}

const Group: React.FC<GroupProps> = ({
  heading,
  open = false,
  updateOpenGroup,
  onNavigate,
  children,
}) => {
  const childrenList = React.Children.toArray(children);

  const [openList, setListOpen] = React.useState(false);

  const anchorRef = React.useRef(null);

  const handleToggle = () => {
    setListOpen(!openList);
  };

  const handleClose = event => {
    if (anchorRef.current && anchorRef.current.contains(event.target)) {
      return;
    }
    setListOpen(false);
  };

  let headingPath = window.location.pathname.replace("/", "").split("/")[0];

  if (headingPath === "ec2") {
    headingPath = "aws";
  }

  const isSelected = headingPath == heading.toLowerCase();

  // n.b. if a Workflow Grouping has no workflows in it don't display it even if
  // it's not explicitly marked as hidden.
  if (childrenList.length === 0) {
    return null;
  }

  return (
    <GroupList
      data-qa="workflowGroup"
      ref={anchorRef}
      aria-controls={openList ? "workflow-options" : undefined}
      aria-haspopup="true"
      onClick={handleToggle}
    >
      <GroupListItem button onClick={() => updateOpenGroup(heading)} selected={isSelected}>
        <GroupHeading align="center">{heading}</GroupHeading>
        <Collapse in={open} timeout="auto" unmountOnExit>
          <Popper
            open={openList}
            anchorEl={anchorRef.current}
            role={undefined}
            transition
            placement="right-start"
            style={{ zIndex: 1200 }}
          >
            <Paper>
              <ClickAwayListener onClickAway={handleClose}>
                <List component="div" disablePadding id="workflow-options">
                  {childrenList.map((c: React.ReactElement) => {
                    return React.cloneElement(c, { onClick: onNavigate });
                  })}
                </List>
              </ClickAwayListener>
            </Paper>
          </Popper>
        </Collapse>
      </GroupListItem>
    </GroupList>
  );
};

interface LinkProps {
  to: string;
  text: string;
  onClick: () => void;
}

const Link: React.FC<LinkProps> = ({ to, text, onClick }) => {
  const isSelected = window.location.pathname.replace("/", "") === to;
  return (
    <LinkListItem
      isSelectedWorkflow={isSelected}
      component={RouterLink}
      onClick={onClick}
      to={to}
      dense
      data-qa="workflowGroupItem"
    >
      <ListItemText>
        <LinkHeading isSelectedWorkflow={isSelected}>{text}</LinkHeading>
      </ListItemText>
    </LinkListItem>
  );
};

interface DrawerProps {
  onClose: () => void;
}

const Drawer: React.FC<DrawerProps> = ({ onClose }) => {
  const { workflows } = useAppContext();
  const [openGroup, setOpenGroup] = React.useState("");

  const updateOpenGroup = (group: string) => {
    if (openGroup === group) {
      setOpenGroup("");
    } else {
      setOpenGroup(group);
    }
  };

  return (
    <DrawerPanel onClose={onClose} data-qa="drawer" variant="permanent">
      {sortedGroupings(workflows).map(grouping => {
        const value = routesByGrouping(workflows)[grouping];
        return (
          <Group
            key={grouping}
            heading={grouping}
            open={openGroup === grouping}
            updateOpenGroup={updateOpenGroup}
            onNavigate={onClose}
          >
            {value.workflows.map(workflow => (
              <Link
                key={workflow.path.replace("/", "")}
                to={workflow.path}
                text={workflow.displayName}
                onClick={onClose}
              />
            ))}
          </Group>
        );
      })}
    </DrawerPanel>
  );
};

export default Drawer;
