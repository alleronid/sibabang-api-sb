module.exports = (sequelize, DataTypes) => {
  const Merchant = sequelize.define('Merchant', {
    merchant_id: {
      type: DataTypes.BIGINT,
      autoIncrement: true,
      primaryKey: true
    },
    merchant_name: {
      type: DataTypes.STRING(191),
      allowNull: false
    },
    status: {
      type: DataTypes.ENUM('ACTIVE', 'DEACTIVE', 'PENDING', ''),
      defaultValue: 'PENDING'
    },
    vendor_id: {
      type: DataTypes.BIGINT,
      allowNull: false,
      defaultValue: 1
    },
    token: {
      type: DataTypes.STRING(191),
      allowNull: false
    },
    env: {
      type: DataTypes.ENUM('PRODUCTION', 'SANDBOX'),
      defaultValue: 'SANDBOX'
    },
    api_key_prod: {
      type: DataTypes.STRING(191)
    },
    cb_key_prod: {
      type: DataTypes.STRING(191)
    },
    api_key_sb: {
      type: DataTypes.STRING(191)
    },
    cb_key_sb: {
      type: DataTypes.STRING(191)
    },
    public_key: {
      type: DataTypes.TEXT
    },
    private_key: {
      type: DataTypes.TEXT
    },
    api_key_ext: {
      type: DataTypes.TEXT
    },
    external_id: {
      type: DataTypes.TEXT
    },
    soundbox: {
      type: DataTypes.TINYINT,
      defaultValue: 0
    },
    nmid: {
      type: DataTypes.STRING(191),
    },
    api_key_sb_vendor: {
      type: DataTypes.STRING(191)
    },
    secret_key_sb_vendor: {
      type: DataTypes.STRING(191)
    },
    api_key_vendor: {
      type: DataTypes.STRING(191)
    },
    secret_key_vendor: {
      type: DataTypes.STRING(191)
    },
    company_id: {
      type: DataTypes.BIGINT,
      allowNull: false
    },
    created_by: {
      type: DataTypes.BIGINT,
    },
    address: {
      type: DataTypes.TEXT
    },
    nmid_qris: {
      type: DataTypes.STRING(191)
    },
    mid_qris: {
      type: DataTypes.STRING(191)
    },
    mid_dana: {
      type: DataTypes.STRING(191)
    },
    mdr_rate: {
      type: DataTypes.DECIMAL(5, 2),
    },
    amount_fee: {
      type: DataTypes.DECIMAL(8, 2),
    },
    createdAt: {
      field: 'created_at',
      type: DataTypes.DATE,
      allowNull: false
    },
    updatedAt: {
      field: 'updated_at',
      type: DataTypes.DATE,
      allowNull: false
    },
    deletedAt : {
      field : 'deleted_at',
      type: DataTypes.DATE,
      allowNull: true
    },
  }, {
    tableName: 'merchants',
    timestamps: true
  });

  return Merchant;
}
