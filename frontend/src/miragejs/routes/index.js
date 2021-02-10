/*
 * Mirage JS guide on Routes: https://miragejs.com/docs/route-handlers/functions
 */

import { Response } from "miragejs";

const WORKSPACE_ID = 1;

export default function routes() {
  // Change this value to simulate response delay.
  // By default development environment has a 400ms delay.
  this.timing = 0;

  this.namespace = "api";

  // User
  this.get("/user/:id");

  // Auth
  this.post("/auth/login", function (schema, request) {
    const loginInfo = this.normalizedRequestAttrs("login-info");
    const user = schema.users.findBy({
      email: loginInfo.username,
      passwordHash: loginInfo.password,
    });
    if (user) {
      return user;
    }
    return new Response(
      401,
      {},
      { errors: loginInfo.username + " not found or incorrect password" }
    );
  });

  this.post("/auth/signup", function (schema, request) {
    const signupInfo = this.normalizedRequestAttrs("signup-info");
    const user = schema.users.findBy({ email: signupInfo.username });
    if (user) {
      return new Response(
        409,
        {},
        { errors: signupInfo.username + " already exists" }
      );
    }
    return schema.users.create(signupInfo);
  });

  // Task
  this.get("/task", function (schema, request) {
    const {
      queryParams: { userid: userId },
    } = request;

    if (userId) {
      return schema.tasks.where((task) => {
        return (
          task.workspaceId == WORKSPACE_ID &&
          (task.creator.id == userId ||
            task.assignee.id == userId ||
            task.subscriberIdList.includes(userId))
        );
      });
    }
    return schema.tasks.none();
  });

  this.get("/task/:id", function (schema, request) {
    const task = schema.tasks.find(request.params.id);
    if (task) {
      return task;
    }
    return new Response(
      404,
      {},
      { errors: "Task " + request.params.id + " not found" }
    );
  });

  this.post("/task", function (schema, request) {
    const ts = Date.now();
    const attrs = this.normalizedRequestAttrs("task");
    const newTask = {
      ...attrs,
      createdTs: ts,
      lastUpdatedTs: ts,
      status: "OPEN",
      workspaceId: WORKSPACE_ID,
    };
    return schema.tasks.create(newTask);
  });

  this.patch("/task/:taskId", function (schema, request) {
    const attrs = this.normalizedRequestAttrs("task-patch");
    return schema.tasks
      .find(request.params.taskId)
      .update({ ...attrs, lastUpdatedTs: Date.now() });
  });

  // environment
  this.get("/environment", function (schema, request) {
    return schema.environments
      .where((environment) => {
        return environment.workspaceId == WORKSPACE_ID;
      })
      .sort((a, b) => a.order - b.order);
  });

  this.post("/environment", function (schema, request) {
    let order = 0;
    const list = schema.environments.where((environment) => {
      return environment.workspaceId == WORKSPACE_ID;
    });
    if (list.length > 0) {
      order = list.sort((a, b) => b.order - a.order).models[0].order + 1;
    }
    const newEnvironment = {
      ...this.normalizedRequestAttrs("environment"),
      workspaceId: WORKSPACE_ID,
      order,
    };
    return schema.environments.create(newEnvironment);
  });

  this.patch("/environment/batch", function (schema, request) {
    const attrs = this.normalizedRequestAttrs("batch-update");
    for (let i = 0; i < attrs.idList.length; i++) {
      const env = schema.environments.find(attrs.idList[i]);
      if (env) {
        const batch = {};
        for (let j = 0; j < attrs.fieldMaskList.length; j++) {
          batch[attrs.fieldMaskList[j]] = attrs.rowValueList[i][j];
        }
        env.update(batch);
      }
    }
    return schema.environments
      .where((environment) => {
        return environment.workspaceId == WORKSPACE_ID;
      })
      .sort((a, b) => a.order - b.order);
  });

  this.patch("/environment/:environmentId", function (schema, request) {
    const attrs = this.normalizedRequestAttrs("environment");
    return schema.environments.find(request.params.environmentId).update(attrs);
  });

  this.delete("/environment/:environmentId", function (schema, request) {
    return schema.environments.find(request.params.environmentId).destroy();
  });

  // Instance
  this.get("/instance", function (schema, request) {
    return schema.instances.where((instance) => {
      return instance.workspaceId == WORKSPACE_ID;
    });
  });

  this.get("/instance/:id", function (schema, request) {
    const instance = schema.instances.find(request.params.id);
    if (instance) {
      return instance;
    }
    return new Response(
      404,
      {},
      { errors: "Instance " + request.params.id + " not found" }
    );
  });

  this.post("/instance", function (schema, request) {
    const newInstance = {
      ...this.normalizedRequestAttrs("instance"),
      workspaceId: WORKSPACE_ID,
    };
    return schema.instances.create(newInstance);
  });

  this.patch("/instance/:instanceId", function (schema, request) {
    const attrs = this.normalizedRequestAttrs("instance");
    return schema.instances.find(request.params.instanceId).update(attrs);
  });

  this.delete("/instance/:instanceId", function (schema, request) {
    return schema.instances.find(request.params.instanceId).destroy();
  });

  // Data Source
  this.get("/instance/:instanceId/datasource", function (schema, request) {
    const instance = schema.instances.find(request.params.instanceId);
    if (instance) {
      return schema.dataSources.where((dataSource) => {
        return dataSource.instanceId == instance.id;
      });
    }
    return new Response(
      404,
      {},
      { errors: "Instance " + request.params.instanceId + " not found" }
    );
  });

  this.get("/instance/:instanceId/datasource/:id", function (schema, request) {
    const instance = schema.instances.find(request.params.instanceId);
    if (instance) {
      const dataSource = schema.dataSources.find(request.params.id);
      if (dataSource) {
        return dataSource;
      }
      return new Response(
        404,
        {},
        { errors: "Data Source " + request.params.id + " not found" }
      );
    }
    return new Response(
      404,
      {},
      { errors: "Instance " + request.params.instanceId + " not found" }
    );
  });

  this.post("/instance/:instanceId/datasource", function (schema, request) {
    const instance = schema.instances.find(request.params.instanceId);
    if (instance) {
      const newDataSource = {
        ...this.normalizedRequestAttrs("data-source"),
        instanceId: request.params.instanceId,
      };
      return schema.dataSources.create(newDataSource);
    }
    return new Response(
      404,
      {},
      { errors: "Instance " + request.params.instanceId + " not found" }
    );
  });

  this.patch(
    "/instance/:instanceId/datasource/:id",
    function (schema, request) {
      const instance = schema.instances.find(request.params.instanceId);
      if (instance) {
        const attrs = this.normalizedRequestAttrs("data-source");
        return schema.dataSources.find(request.params.id).update(attrs);
      }
      return new Response(
        404,
        {},
        { errors: "Instance " + request.params.instanceId + " not found" }
      );
    }
  );

  this.delete(
    "/instance/:instanceId/datasource/:id",
    function (schema, request) {
      const instance = schema.instances.find(request.params.instanceId);
      if (instance) {
        return schema.dataSources.find(request.params.id).destroy();
      }
      return new Response(
        404,
        {},
        { errors: "Instance " + request.params.instanceId + " not found" }
      );
    }
  );

  // Bookmark
  this.get("/bookmark", function (schema, request) {
    return schema.bookmarks.where({
      workspaceId: WORKSPACE_ID,
    });
  });

  // Activity
  this.get("/activity", function (schema, request) {
    return schema.activities.where({
      workspaceId: WORKSPACE_ID,
    });
  });

  // Group
  this.get("/group", function (schema, request) {
    const {
      queryParams: { userid: userId },
    } = request;

    if (userId) {
      return schema.groups.where((group) => {
        return (
          group.workspaceId == WORKSPACE_ID &&
          group.groupRoleIds?.find((roleId) => {
            return schema.groupRoles.where({
              userId,
              groupId: group.id,
            });
          })
        );
      });
    }
    return schema.groups.where({
      workspaceId: WORKSPACE_ID,
    });
  });

  // Project
  this.get("/project", function (schema, request) {
    const {
      queryParams: { groupid: groupId, userid: userId },
    } = request;

    if (groupId && userId) {
      return new Response(
        400,
        {},
        { errors: "groupid and userid can't be specified together" }
      );
    }

    if (!groupId && !userId) {
      return schema.projects.where({
        workspaceId: WORKSPACE_ID,
      });
    }

    const groupIdList = [];
    if (groupId) {
      groupIdList.push(groupId);
    } else if (userId) {
      const groupList = schema.groups.where((group) => {
        return (
          group.workspaceId == WORKSPACE_ID &&
          group.groupRoleIds?.find((roleId) => {
            return schema.groupRoles.where({
              userId,
              groupId: group.id,
            });
          })
        );
      });
      groupList.models.forEach((group) => {
        groupIdList.push(group.id);
      });
    }

    if (groupIdList.length) {
      return schema.projects.where((project) => {
        return (
          project.workspaceId == WORKSPACE_ID &&
          groupIdList.includes(project.groupId)
        );
      });
    }
    return schema.environments.none();
  });

  this.get("/project/:id/task", function (schema, request) {
    const project = schema.projects.find(request.params.id);
    if (project) {
      return schema.tasks.where({
        projectId: project.id,
      });
    }
    return schema.environments.none();
  });

  this.get("/project/:id/task/:taskId", function (schema, request) {
    const project = schema.projects.find(request.params.id);
    if (project) {
      return schema.tasks.findBy({
        slug: request.params.taskId,
      });
    }
    return null;
  });

  this.get("/project/:id/repository", function (schema, request) {
    const project = schema.projects.find(request.params.id);
    if (project) {
      return schema.repositories.findBy({
        projectId: project.id,
      });
    }
    return null;
  });
}
