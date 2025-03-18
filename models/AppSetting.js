module.exports = (sequelize, DataTypes) => {
  const AppSetting = sequelize.define('AppSetting', {
    key: {
      type: DataTypes.STRING(191),
      primaryKey: true, // Huruf "P" besar
    },
    value: {
      type: DataTypes.STRING(191),
    },
    is_active: {
      type: DataTypes.TINYINT,
      defaultValue: 1,
    }
  }, {
    tableName: "app_settings",
    timestamps: false, // Menghilangkan createdAt & updatedAt
  });

  return AppSetting;
};
